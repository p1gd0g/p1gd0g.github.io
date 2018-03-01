# Linkable Ring Signature with Unconditional Anonymity

## Sign

1. Compute $e=H(event)$ and $t=e^x$.

1. Randomly pick $r_x, r_y, c_1 ,..., c_{s-1},c_{s+1},...,c_n \in_R \mathbb{Z}_p$ and compute $K=g^{r_x}h^{r_y}\prod_{i=1,i\neq s}^nZ_i^{c_i}$, $K'=e^{r_x}t^{\sum_{i=1,i\neq s}^nc_i}$.

1. Find $c_s$ such that $c_1+...+c_n$ mod $p=H'(\mathcal{Y}||event||t||M||K||K')$.

1. Compute $\tilde{x}=r_x-c_sx$ mod $p$, $\tilde{y}=r_y-c_s$ mod $p$.