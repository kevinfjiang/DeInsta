# EtheReddit
> **Decentralized Reddit clone built on Ethereum smart contracts**


![License](https://img.shields.io/github/license/kevinfjiang/BirthdaySlackbot.svg)![Lcommit](https://img.shields.io/github/last-commit/kevinfjiang/EtheReddit)


### **Premise:**
EtheReddit takes the blog post feature, comments, and upvote algorithim of Reddit and deploys it on the blockchain to protect against censorship. Users can also send posters Ethereum if they like the post. Will consider expanding the character limit to full out medium style blogs and rendering html and everything.


### **Tech/Design choices:**

**Go**: Favorite backend language right now, Geth autoo encodes targets from Solidity -> Go

**Solidity**: Languagge of Ethereum smart contracts

**Ethereum**: Language has really robust smart contract support so it's super easy to set up Posts as contracts where you can input money into them. 


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
- [ ] Develop the algorithim for top active posts

**Frontend**
- [ ] Develop front end that is passable (Not a front end guy lol)
</details>