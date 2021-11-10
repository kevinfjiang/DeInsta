## **DeInsta**
> **Decentralized Insta clone built on Ethereum smart contracts and IPFS file share**


![License](https://img.shields.io/github/license/kevinfjiang/DeInsta.svg) ![Lcommit](https://img.shields.io/github/last-commit/kevinfjiang/DeInsta)


### **Premise:**
DeInsta takes the image sharing feature of instagram and adds the storage capabilities on the blockchain so users have more control of what they can post and share. It also allows for direct exchange of photos through IPFS, a peer-to-peer file share system. Intended as a proof of concept and testing of decentralized platforms.
  
=======


### **Tech/Design choices:**

**Go**: Favorite backend language right now, Geth auto encodes targets from Solidity -> Go

**Solidity**: Language of Ethereum smart contracts

**Ethereum**: Language has smart contracts

**React**: UI, pretty self explanatory

**IPFS**: Peer to peer file share, much cheaper than storing on the blockchain, which is more secure but crazy expensive. This way, we ensure that posting is cheap-ish


### **TODO:**
#### Personal stuff for Kevin to organize and show what's currently accocmplished in the project
<details>
<summary>Project TODOs</summary>
<br>

**Blockchain/DB**
- [x] Write Contract Solidity code
- [x] Enable deletes and write the registry
- [x] Enable votes, posts, comments
- [x] Finish writing Solidity code
- [ ] Expand character limit, consider allowing html


**Backend**
- [x] Create bindings for Solidity to Go with Geth
- [x] Enable the blockchain to emit events
- [ ] Develop the backend for the interactions between posts
- [ ] Develop the algorithm for top active posts


**Frontend**
- [ ] Develop front end that is passable (Not a front end guy lol)
- [ ] Choose a frontend from existing builds and modify the API to be workable with golang
</details>