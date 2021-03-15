#### TANGLE
An append only data structure where each item references two other items.

#### MESSAGE
The object type which is gossiped between neighbors. All gossiped information is included in a message.

#### MESSAGE TANGLE
The collection of all messages.

#### OBJECT
The most basic unit of information of the IOTA Protocol. Each object has a type and size and contains data.

#### CORE OBJECT TYPE
An object type which must be parsed by all users. Core object types include:
* Value objects
* FPC Opinion objects
* DRNG objects
* Salt declaration objects
* Generic data objects

#### GENERIC DATA OBJECT
The most basic object type.

#### PAYLOAD
A field in an object which can only be filled by another object.

#### VALUE TANGLE
The collection of all value objects.

#### VALUE TRANSFER APPLICATION
The application which maintains the ledger state.

#### VALUE OBJECT
The basic object of the value transfer application.

#### TRANSACTION
The payload of a value object. It contains the particulars of a transfer of funds.

#### UTXO
Unspent transaction output.

#### TIP
A message/value object that has not yet been approved.

#### TIP SELECTION
The process of selecting previous messages/value objects to be referenced by a new message/value object. These references are where a message/value object attaches to the existing data structure. IOTA only enforces that a message/value object approves two other messages/value objects, but the tip selection strategy is left up to the user (with a good default provided by IOTA).

#### NETWORK LAYER
This layer manages the lower layers of internet communication like TCP. It is the most technical, and in some ways the least interesting. In this layer, the connections between nodes are managed by the autopeering and peer discovery modules and the gossip protocol.

#### COMMUNICATION LAYER
This layer stores and communicates information. This layer contains the “distributed ledger” or the tangle. The rate control and timestamps are in this layer too.

#### APPLICATION LAYER
The IOTA Protocol allows for a host of applications to run on the message tangle. Anybody can design an application, and users can decide which applications to run on their nodes. These applications will all use the communication layer to broadcast and store data.

#### CORE APPLICATIONS
Applications that are necessary for the protocol to operate. These include for example:
* The value transfer application
* The distributed random number generator (DRNG for short)
* The Fast Probabilistic Consensus (FPC) protocol

#### FAUCET
A test application issuing funds on request.

#### BLOCKCHAIN BOTTLENECK
As more transactions are issued, the block rate and size become a bottleneck in the system. It can no longer include all incoming transactions promptly. Attempts to speed up block rates will introduce more orphan blocks (blocks being left behind) and reduce the security of the blockchain.

#### CONSENSUS
Agreement on a specific datum or value in distributed multi-agent systems, in the presence of faulty processes.

#### COORDINATOR
A trusted entity that issues milestones to guarantee finality and protect the Tangle against attacks.

#### ECLIPSE ATTACK
A cyber-attack that aims to isolate and attack a specific user, rather than the whole network.

#### FINALITY
The property that once a transaction is completed there is no way to revert or alter it. This is the moment when the parties involved in a transfer can consider the deal done. Finality can be deterministic or probabilistic.

#### HISTORY
The list of transactions directly or indirectly approved by a given transaction.

#### LOCAL MODIFIERS
Custom conditions that nodes can take into account during tip selection. In IOTA, nodes do not necessarily have the same view of the Tangle; various kinds of information only locally available to them can be used to strengthen security.

#### MANA
The reputation of a node is based on a virtual token called mana. This reputation, working as a Sybil protection mechanism, is important for issuing more transactions (see Module 3) and having a higher influence during the voting process (see Module 5).

#### MESSAGE OVERHEAD
The additional information (metadata) that needs to be sent along with the actual information (data). This can contain signatures, voting, heartbeat signals, and anything that is transmitted over the network but is not the transaction itself.

#### MILESTONES
Milestones are transactions signed and issued by the Coordinator. Their main goal is to help the Tangle to grow healthily and to guarantee finality. When milestones directly or indirectly approve a transaction in the Tangle, nodes mark the state of that transaction and its entire history as confirmed.

#### MINING RACES
In PoW-based DLTs, competition between nodes to obtain mining rewards and transaction fees are known as mining races. These are undesirable as they favor more powerful nodes, especially those with highly optimized hardware like ASICs. As such, they block participation by regular or IoT hardware and are harmful for the environment.

#### NAKAMOTO CONSENSUS
Named after the originator of Bitcoin, Satoshi Nakamoto, Nakamoto consensus describes the replacement of voting/communication between known agents with a cryptographic puzzle (Proof-of-Work). Completing the puzzle determines which agent is the next to act.

#### NEIGHBORS
Network nodes that are directly connected and can exchange messages without intermediate nodes.

#### NODE
A machine which is part of the IOTA network. Its role is to issue new transactions and to validate existing ones.

#### ORPHAN
A transaction (or block) that is not referenced by any succeeding transaction (or block). An orphan is not considered confirmed and will not be part of the consensus.

#### PARASITE-CHAIN ATTACKS
A double spend attack on the Tangle. Here an attacker attempts to undo a transaction by building an alternative Tangle in which the funds were not spent. They then try to get the majority of the network to accept the alternative Tangle as the legitimate one.

#### PEERING
The procedure of discovering and connecting to other network nodes.

#### PROOF-OF-WORK
Data which is difficult (costly, time-consuming) to produce but easy for others to verify.

#### RANDOM WALK
A mathematical object that describes a path, which consists of a succession of random steps in some mathematical space.

#### REATTACHMENT
Resending a transaction by redoing tip selection and referencing newer tips by redoing PoW.

#### SMALL-WORLD NETWORK
A network in which most nodes can be reached from every other node by a small number of intermediate steps.

#### SOLIDIFICATION TIME
The solidification time is the point at which the entire history of a transaction has been received by a node.

#### SPLITTING ATTACKS
An attack in which a malicious node attempts to split the Tangle into two branches. As one of the branches grows the attacker publishes transactions on the other branch to keep both alive.Splitting attacks attempt to slow down the consensus process or conduct a double spend.

#### SUBTANGLE
A consistent section of the Tangle (i.e. a subset of messages/value objects), such that each included message/value object also includes its referenced messages/value objects.

#### SYBIL ATTACK
An attempt to gain control over a peer-to-peer network by forging multiple fake identities.