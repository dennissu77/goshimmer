package epochs

import (
	"fmt"
	"math"
	"sync"

	"github.com/iotaledger/hive.go/byteutils"
	"github.com/iotaledger/hive.go/cerrors"
	"github.com/iotaledger/hive.go/identity"
	"github.com/iotaledger/hive.go/marshalutil"
	"github.com/iotaledger/hive.go/objectstorage"
	"github.com/iotaledger/hive.go/stringify"
	"golang.org/x/xerrors"
)

// region EpochID /////////////////////////////////////////////////////////////////////////////////////////////////

// ID is the Epoch's ID.
type ID uint64

// IDFromBytes unmarshals an ID from a sequence of bytes.
func IDFromBytes(bytes []byte) (id ID, consumedBytes int, err error) {
	marshalUtil := marshalutil.New(bytes)
	if id, err = IDFromMarshalUtil(marshalUtil); err != nil {
		err = xerrors.Errorf("failed to parse ID from MarshalUtil: %w", err)
		return
	}
	consumedBytes = marshalUtil.ReadOffset()

	return
}

// IDFromMarshalUtil unmarshals an ID using a MarshalUtil (for easier unmarshaling).
func IDFromMarshalUtil(marshalUtil *marshalutil.MarshalUtil) (id ID, err error) {
	untypedID, err := marshalUtil.ReadUint64()
	if err != nil {
		err = xerrors.Errorf("failed to parse ID (%v): %w", err, cerrors.ErrParseBytesFailed)
		return
	}
	id = ID(untypedID)
	return
}

// Bytes returns a marshaled version of the ID.
func (i ID) Bytes() []byte {
	return marshalutil.New(marshalutil.Uint64Size).WriteUint64(uint64(i)).Bytes()
}

// String returns a human readable version of the ID.
func (i ID) String() string {
	return fmt.Sprintf("EpochID(%X)", uint64(i))
}

// endregion ///////////////////////////////////////////////////////////////////////////////////////////////////////////

// region Epoch /////////////////////////////////////////////////////////////////////////////////////////////////////

type Epoch struct {
	objectstorage.StorableObjectFlags

	id   ID
	mana map[identity.ID]float64

	manaMutex sync.RWMutex
}

func New(id ID) *Epoch {
	return &Epoch{
		id:   id,
		mana: make(map[identity.ID]float64),
	}
}

// EpochFromBytes parses the given bytes into an Epoch.
func EpochFromBytes(bytes []byte) (result *Epoch, consumedBytes int, err error) {
	marshalUtil := marshalutil.New(bytes)
	result, err = EpochFromMarshalUtil(marshalUtil)
	consumedBytes = marshalUtil.ReadOffset()
	return
}

// EpochFromMarshalUtil parses a new Epoch from the given marshal util.
func EpochFromMarshalUtil(marshalUtil *marshalutil.MarshalUtil) (result *Epoch, err error) {
	result = &Epoch{
		mana: make(map[identity.ID]float64),
	}
	if result.id, err = IDFromMarshalUtil(marshalUtil); err != nil {
		err = xerrors.Errorf("failed to parse EpochID from MarshalUtil: %w", err)
		return
	}

	var nodesCount uint32
	if nodesCount, err = marshalUtil.ReadUint32(); err != nil {
		err = xerrors.Errorf("failed to parse nodes count from MarshalUtil: %w", err)
		return
	}

	for i := 0; i < int(nodesCount); i++ {
		var nodeID identity.ID
		if nodeID, err = identity.IDFromMarshalUtil(marshalUtil); err != nil {
			err = xerrors.Errorf("failed to parse nodeID from MarshalUtil: %w", err)
			return
		}

		var mana uint64
		if mana, err = marshalUtil.ReadUint64(); err != nil {
			err = xerrors.Errorf("failed to parse mana value from MarshalUtil: %w", err)
			return
		}

		result.mana[nodeID] = math.Float64frombits(mana)
	}

	return
}

// EpochFromObjectStorage is the factory method for Epoch stored in the ObjectStorage.
func EpochFromObjectStorage(key []byte, data []byte) (result objectstorage.StorableObject, err error) {
	if result, _, err = EpochFromBytes(byteutils.ConcatBytes(key, data)); err != nil {
		err = xerrors.Errorf("failed to parse Epoch from bytes: %w", err)
		return
	}

	return
}

func (e *Epoch) ID() ID {
	return e.id
}

func (e *Epoch) AddNode(id identity.ID) {
	e.manaMutex.Lock()
	defer e.manaMutex.Unlock()

	e.mana[id] = 0
}

func (e *Epoch) Mana() (mana map[identity.ID]float64) {
	e.manaMutex.RLock()
	defer e.manaMutex.RUnlock()

	mana = make(map[identity.ID]float64)
	for nodeID, m := range e.mana {
		mana[nodeID] = m
	}
	return
}

func (e *Epoch) TotalMana() float64 {
	e.manaMutex.RLock()
	defer e.manaMutex.RUnlock()

	var total float64
	for _, mana := range e.mana {
		total += mana
	}
	return total
}

// Bytes returns a marshaled version of the whole Epoch object.
func (e *Epoch) Bytes() []byte {
	return byteutils.ConcatBytes(e.ObjectStorageKey(), e.ObjectStorageValue())
}

// ObjectStorageKey returns the key of the stored Epoch object.
// This returns the bytes of the ID.
func (e *Epoch) ObjectStorageKey() []byte {
	return e.id.Bytes()
}

// ObjectStorageValue returns the value of the stored Epoch object.
func (e *Epoch) ObjectStorageValue() []byte {
	marshalUtil := marshalutil.New()

	e.manaMutex.RLock()
	defer e.manaMutex.RUnlock()

	marshalUtil.WriteUint32(uint32(len(e.mana)))
	for nodeID, mana := range e.mana {
		marshalUtil.Write(nodeID)
		marshalUtil.WriteUint64(math.Float64bits(mana))
	}

	return marshalUtil.Bytes()
}

// Update updates the Epoch.
// This should never happen and will panic if attempted.
func (e *Epoch) Update(objectstorage.StorableObject) {
	panic("updates disabled")
}

// String returns a human readable version of the Epoch.
func (e *Epoch) String() string {
	builder := stringify.StructBuilder("Epoch", stringify.StructField("ID", e.id))

	e.manaMutex.RLock()
	for nodeID, mana := range e.mana {
		builder.AddField(stringify.StructField(nodeID.String(), fmt.Sprintf("%.2f", mana)))
	}
	e.manaMutex.RUnlock()
	builder.AddField(stringify.StructField("TotalMana", fmt.Sprintf("%.2f", e.TotalMana())))
	return builder.String()
}

// interface contract (allow the compiler to check if the implementation has all of the required methods).
var _ objectstorage.StorableObject = &Epoch{}

// endregion ///////////////////////////////////////////////////////////////////////////////////////////////////////////

// region CachedEpoch ///////////////////////////////////////////////////////////////////////////////////////////////

// CachedEpoch is a wrapper for a stored cached object representing an Epoch.
type CachedEpoch struct {
	objectstorage.CachedObject
}

// Unwrap unwraps the CachedEpoch into the underlying Epoch.
// If stored object cannot be cast into an Epoch or has been deleted, it returns nil.
func (c *CachedEpoch) Unwrap() *Epoch {
	untypedObject := c.Get()
	if untypedObject == nil {
		return nil
	}

	typedObject := untypedObject.(*Epoch)
	if typedObject == nil || typedObject.IsDeleted() {
		return nil
	}

	return typedObject

}

// Consume consumes the CachedEpoch.
// It releases the object when the callback is done.
// It returns true if the callback was called.
func (c *CachedEpoch) Consume(consumer func(epoch *Epoch), forceRelease ...bool) (consumed bool) {
	return c.CachedObject.Consume(func(object objectstorage.StorableObject) {
		consumer(object.(*Epoch))
	}, forceRelease...)
}

// endregion ///////////////////////////////////////////////////////////////////////////////////////////////////////////