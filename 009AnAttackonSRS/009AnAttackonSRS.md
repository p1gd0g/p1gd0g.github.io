# An Attack on SRS in the standard model

Thanks for your advises!

---

First we recall your attack.

Assume we have a ring signature:

$$e(A_1,T_1)e(A_2,T_2)=e(P,g_2)$$

Then we attack two public keys $(A_3,A_4)$.

$$A_1'=A_1A_4/A_3$$

$$A_2'=A_2A_4/A_3$$

$$e(A_1'A_3/A_4,T_1)e(A_2'A_3/A_4,T_2)=e(P,g_2)$$

$$e(A_1',T_1)e(A_2',T_2)e(A_3,T_1T_2)e(A_4,1/T_1T_2)=e(P,g_2)$$

We output a "valid" signature: $(P,T_1,T_2,T_3=T_1T_2,T_4=1/T_1T_2)$.

---

However, as we argued in the overview.

> To make their ring signature scheme independent with the \textsf{crs}, they divide the \textsf{crs} of \textsf{NIZK} into a part of each verification key, achieving that the \textsf{crs} of \textsf{NIZK} is not the \textsf{crs} of ring signature.

which means the crs is different in different ring signature (actually, there is not a crs in the ring signature). It follows that you couldn't denote some member of a ring only by $A$, which is $z'/z_i$ in the first NIZK of the ring signature, where $z'$ is the re-randomized signing key and $z_i$ is one element of the verification key $(z,k,C)$.

At the same time, we note that

$$T=\prod C_i$$

which means the real crs in the NIZK is defined by every member of the ring. Therefore, your ring signature couldn't be verified.

Assume you find correct $C_i$ with negligible probability, you could not generate the correct $y$.

In summary, you should consider the whole ring signatuer but not the NIZK arguments of knowledge only. The NIZK is independent as well as its security.