# Bulletproofs

## Linear Range Proof

First we give the relation of range proof.

$$\{(g,h\in\mathbb{G},V,n;v,\gamma\in\mathbb{Z}_p):V=h^\gamma g^v\wedge v\in[0,2^n-1]\}$$

Decode the value $v$ into $\mathbf{a}_L\in\{0,1\}^n$ s.t. $\langle\mathbf{a}_L,\mathbf{2}^n\rangle=v$ and $\mathbf{a}_R=\mathbf{a}_L-\mathbf{1}^n$.

For a random $y\in\mathbb{Z}_p$ from verifier, the prover should prove that

$$\langle\mathbf{a}_L,\mathbf{2}^n\rangle=v\wedge\langle\mathbf{a}_L,\mathbf{a}_R\circ\mathbf{y}^n\rangle=0\wedge\langle\mathbf{a}_L-\mathbf{1}^n-\mathbf{a}_R,\mathbf{y}^n\rangle=0$$

For a random $z\in\mathbb{Z}_p$, it follows that

$$z^2\cdot\langle\mathbf{a}_L,\mathbf{2}^n\rangle+z\cdot\langle\mathbf{a}_L-\mathbf{1}^n-\mathbf{a}_R,\mathbf{y}^n\rangle+\langle\mathbf{a}_L,\mathbf{a}_R\circ\mathbf{y}^n\rangle=z^2\cdot v$$

$$\langle\mathbf{a}_L-z\cdot\mathbf{1}^n,\mathbf{y}^n\circ(\mathbf{a}_R+z\cdot\mathbf{1}^n)+z^2\cdot\mathbf{2}^n=z^2\cdot v+\delta(y,z)$$

where $\delta(y,z)=(z-z^2)\cdot\langle\mathbf{1}^n,\mathbf{y}^n\rangle-z^3\langle\mathbf{1}^n,\mathbf{2}^n\rangle\in\mathbb{Z}_p$.

Prover picks randomly $\alpha,\rho\leftarrow\mathbb{Z}_p$ and $\mathbf{s}_L,\mathbf{s}_R\leftarrow\mathbb{Z}_p^n$, computes

$$A=h^\alpha\mathbf{g}^{\mathbf{a}_L}\mathbf{h}^{\mathbf{a}_R}$$

$$S=h^\rho\mathbf{g}^{\mathbf{s}_L}\mathbf{h}^{\mathbf{s}_R}$$

Verifier receives $A$ and $S$, sends $y$ and $z$ to prover.

Prover picks $\tau_1,\tau_2\leftarrow\mathbb{Z}_p$ and computes $T_i=g^{t_i}h^{\tau_i}$, sends $T_i$ to verifier.

Verifier picks $x\leftarrow\mathbb{Z}_p^\star$ and sends it to prover.

Prover computes

$$\mathbf{l}=l(x)=\mathbf{a}_L-z\cdot\mathbf{1}^n+\mathbf{s}_L\cdot x$$

$$\mathbf{r}=r(x)=\mathbf{y}^n\circ(\mathbf{a}_R+z\cdot\mathbf{1}^n+\mathbf{s}_R\cdot x)+z^2\cdot\mathbf{2}^n$$

$$\hat{t}=\langle\mathbf{l},\mathbf{r}\rangle$$

$$\tau_x=\tau_2\cdot x^2+\tau_1\cdot x+z^2\cdot\gamma$$

$$\mu=\alpha+\rho\cdot x$$

where $\mathbf{s}_*\in\mathbb{Z}_p$ are used to blind $\mathbf{a}_*$. Pvover sends $\tau_x,\mu,\hat{t},\mathbf{l},\mathbf{r}$ to verifier.

Verifier checks

$$h_i'=h_i^{(y^{-1}+1)}$$

$$g^{\hat{t}}h^{\tau_x}\overset{?}{=}V^{z^2}\cdot g^{\delta(y,z)}\cdot T_1^x\cdot T_2^{x^2}$$

$$P=A\cdot S^x\cdot\mathbf{g}^{-z}\cdot(\mathbf{h}')^{z\cdot\mathbf{y}^n+z^2\cdot\mathbf{2}^n}$$

$$P\overset{?}{=}h^\mu\cdot\mathbf{g}^{\mathbf{l}}\cdot(\mathbf{h}')^{\mathbf{r}}$$

$$\hat{t}\overset{?}{=}\langle\mathbf{l},\mathbf{r}\rangle$$