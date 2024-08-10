## Total ordered multicast for distributed events

- 3 Nodes - Distributefd System from the outside it's a single system
- The system would have to keep track of multiple "accounts" and their balances. And periodically process `transactions` where money is moved between `accounts`.
- Within the system - if one node initiates or processes the transaction - it has to be done by everyone else as well
- But it has to be done in a way that all nodes process the same transactions in the same order - `total ordered multicast - ISIS`
- 