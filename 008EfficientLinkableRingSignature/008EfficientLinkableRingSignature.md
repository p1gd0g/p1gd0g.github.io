# Efficient Linkable Ring Signature

## Sign

1. Compute $I=x_sH_p(G)$.

1. Randomly pick $\alpha,c_1,...,c_{s-1},c_{s+1},...,c_n$ and compute:

    1. $L=\alpha G+\sum_{i\neq s}^{n}{c_i}P_i$.

    1. $R=\alpha H_p(G)+\sum_{i\neq s}^{n}c_iI$.

1. Compute $c_s=H_s(L,R,M)-\sum_{i\neq s}^nc_i$.

1. Compute $r=\alpha -c_sx_s$.

1. Output the signature $\sigma=(I,r,c_1,...,c_n)$.

## Verify

1. Compute $L=rG+\sum_i^nc_iP_i$ and $R=rH_p(G)+\sum_i^nc_iI$.

1. Check whether $\sum_i^nc_i=H_s(L,R,M)$.