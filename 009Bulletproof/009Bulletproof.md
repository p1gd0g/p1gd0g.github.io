# Bulletproof

## Linear Range Proof

First we give the relation of range proof.

$$\{(g,h\in\mathbb{G},V,n;v,\gamma\in\mathbb{Z}_p):V=h^\gamma g^v\wedge v\in[0,2^n-1]\}$$

Decode the value $v$ into $\mathbf{a}_L\in\{0,1\}^n$ s.t. $\left\langle\mathbf{a}_L,\mathbf{2}^n\right\langle=v$ and $\mathbf{a}_L+\mathbf{a}_R=\mathbf{1}^n$.

For a random $y\in\mathbb{Z}_p$ from verifier, the prover should prove that

$$<\mathbf{a}_L,\mathbf{2}^n>$$