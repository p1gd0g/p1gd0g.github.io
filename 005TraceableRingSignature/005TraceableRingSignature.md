# Traceable Ring Signature

## Sign

To sign message $m\in \left\lbrace 0,1\right\rbrace^*$ with respect to tag $L=(issue,pk_N)$, using the secret-key $sk_s$, preceed as follows:

1. Compute $h=H(L)$ and $\sigma_s=h^{x_s}$, using $x_s\in \mathbb{Z}_q$.

1. Set $A_0=H'(L,m)$ and $A_1=(\frac{\sigma_s}{A_0})^{1/s}$.

1. For all $i\neq s$, compute $\sigma_i=A_0A_1^i\in G$.

1. Generate signature $(r_N,c_N)$ on $(L,m)$ as follows:

    1. Pick up random $\alpha\leftarrow \mathbb{Z}_q$ and set $L_s=g^{\alpha}, R_s=h^\alpha\in G$.

    1. Pick up at random $r_i,c_i\leftarrow \mathbb{Z}_q$, and set $L_i=g^{r_i}y_i^{c_i}, R_i=h^{r_i}\sigma_i^{c_i}\in G$ for every $i\neq s$.

    1. Set $c=H^{\prime\prime}(L,A_0,A_1,L_N,R_N)$.

    1. Set $c_s=c-\sum_{i\neq s}c_i$ and $r_s=\alpha-c_sx_s$.

1. Output $\sigma=(A_1,c_N,r_N)$ as the signature on $(L,m)$.

## Verify

To verify signature $\sigma=(A_1,c_N,r_N)$ on message $m$ with respect to tag $L$, check the following:

1. Parse $L$ as $(issue,pk_N)$. Set $h=H(L)$ and $A_0=H'(L,m)$, and compute $\sigma_i=A_0A_1^i\in G$ for all $i\in N$.

1. Compute $L_i=g^{r_i}y_i^{c_i}$ and $R_i=h^{r_i}\sigma_i^{c_i}$ for all $i\in N$.

1. Check that $H^{\prime\prime}(L,m,A_0,A_1,L_N,R_N)\equiv \sum_{i\in N}c_i$.

## Trace

To check the relation between $(m,\sigma)$ and $(m',\sigma')$, with respect to the same tag $L$ where $\sigma=(A_1,c_N,r_N)$ and $\sigma'=(A_1',c_N',r_N')$, check the following:

1. Parse $L$ as $(issue,pk_N)$. Set $h=(L)$ and $A_0=H'(L,m)$, and compute $\sigma_i=A_0A_1^i\in G$ for all $(i\in N)$.

1. For all $i\in N$, if $\sigma_i=\sigma_i'$, store $pk_i$ in $\textbf{List}$.

1. Output $pk$ if $pk$ is the only entry in $\textbf{List}$; "linked" else if $\textbf{List}=pk_N$;; "indep" otherwise.