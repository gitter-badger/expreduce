possibleExponents[n_Integer, m_Integer] := 
 Flatten[Permutations /@ ((PadRight[#, m]) & /@ 
     IntegerPartitions[n, m]), 1];
genVars[addends_List, exponents_List] := 
 Product[addends[[ExpandUnique`i]]^
   exponents[[ExpandUnique`i]], {ExpandUnique`i, 1, Length[addends]}];
genExpand[addends_List, exponents_List] := 
 Sum[(Multinomial @@ exponents[[ExpandUnique`i]])*
   genVars[addends, exponents[[ExpandUnique`i]]], {ExpandUnique`i, 1, 
   Length[exponents]}];
Expand[a_] := a //. {
    s_Plus^n_Integer?Positive :> 
     genExpand[List @@ s, possibleExponents[n, Length[s]]],
    c_*s_Plus :> ((c*#) &) /@ s
    };

PolynomialQ[p_Plus, v_] :=
  AllTrue[List @@ p, (PolynomialQ[#, v]) &];
PolynomialQ[p_.*v_^Optional[exp_Integer], v_] :=
  If[FreeQ[p, v] && Positive[exp], True, False];
PolynomialQ[p_, v_] := If[FreeQ[p, v], True, False];

Exponent[expr_/p_Plus,var_,head_]:=Exponent[expr,var,head];
Exponent[expr_,var_,head_]:=Module[{e=expr,v=var,h=head,theCases,toCheck},
    toCheck=expr//Expand;
    toCheck=If[Head[toCheck]===Plus,toCheck,{toCheck}];
    theCases=Cases[toCheck,p_.*v^Optional[exp_]->exp]//DeleteDuplicates;
    If[Length[theCases]=!=Length[toCheck],PrependTo[theCases,0]];
    h@@theCases
];
Exponent[expr_,var_]:=Exponent[expr,var,Max];

Attributes[Exponent] = {Listable, Protected};

ExpreduceSingleCoefficient[inP_, inTerm_] :=
  Module[{p = inP, term = inTerm, pat},
   (*If[MatchQ[p,term],Return[1]];*)
   pat = If[term === 1, a_?NumberQ, Optional[a_]*term];
   (*pat=Optional[a_]*term;*)
   If[MatchQ[p, pat],
    (p) /. pat -> a, 0]
   ];
Coefficient[p_, var_, exp_] := Coefficient[p, var^exp];
Coefficient[inP_, inTerm_] :=
  Module[{p = inP, term = inTerm, toMatch},
   toMatch = p // Expand;
   If[Head[toMatch] === Plus,
    Map[ExpreduceSingleCoefficient[#, term] &, toMatch],
    ExpreduceSingleCoefficient[toMatch, term]]
   ];

Attributes[Coefficient] = {Listable, Protected};

ExpreduceLeadingCoeff[p_, x_] := Coefficient[p, x^Exponent[p, x]];
PolynomialQuotientRemainder[inp_, inq_, v_] :=
  Module[{a = inp, b = inq, x = v, r, d, c, i, s, q},
   q = 0;
   r = a;
   d = Exponent[b, x];
   c = ExpreduceLeadingCoeff[b, x];
   i = 1;
   While[Exponent[r, x] >= d && i < 20,
    s = (ExpreduceLeadingCoeff[r, x]/c)*x^(Exponent[r, x] - d);
    q = q + s;
    r = r - s*b;
    i = i + 1;
    ];
   {q, r} // Expand
   ];
Attributes[PolynomialQuotientRemainder] = {Protected};
PolynomialQuotient[inp_, inq_, v_] :=
  PolynomialQuotientRemainder[inp, inq, v][[1]];
Attributes[PolynomialQuotient] = {Protected};
PolynomialRemainder[inp_, inq_, v_] :=
  PolynomialQuotientRemainder[inp, inq, v][[2]];
Attributes[PolynomialRemainder] = {Protected};

ExpreduceConstantTerm[c_?NumberQ] := {c, 1};
ExpreduceConstantTerm[c_?NumberQ*e_] := {c, e};
ExpreduceConstantTerm[e_] := {1, e};
FactorTermsList[expr_] := Module[{e = expr, toFactor, cTerms, c},
   toFactor = e // Expand;
   If[Head[toFactor] =!= Plus,
    Return[ExpreduceConstantTerm[toFactor]]
    ];
   (* Parens are necessary due to precedence issue. *)
   cTerms = ((ExpreduceConstantTerm /@ (List @@ toFactor)) // 
       Transpose)[[1]];
   c = GCD @@ cTerms;
   If[Last[cTerms] < 0, c = -c];
   {c, toFactor/c // Expand}
   ];
Attributes[FactorTermsList] = {Protected};

(*PolySubresultantGCD[inA_, inB_, inX_] := 
  Module[{u = inA, v = inB, x = inX, h, delta, beta, newU, newV, i},
   Print[u];
   Print[v];
   Print[x];
   h = 1;
   i = 1;
   While[v =!= 0 && i < 20,
    delta = Exponent[u, x] - Exponent[v, x];
    Print[delta];
    beta = (-1)^(delta + 1)*Exponent[u, x]*h^delta;
    Print[beta];
    h = h*(Exponent[v, x]/h)^delta;
    Print[h];
    newU = v;
    newV = PolynomialRemainder[u, v, x]/beta;
    Print[newV];
    u = newU;
    v = newV;
    i = i + 1;
    ];
   If[Exponent[u, x] == 0, 1, u]
   ];*)
(* doesn't work with rational functions yet. *)

(* Looks like prefactored inputs remain factored. *)
(*PolynomialGCD[inA_, inB_] := 
  FactorTermsList[
    PolySubresultantGCD[inA, inB, Variables[inA][[1]]]][[2]];

Attributes[PolynomialGCD] = {Listable, Protected};*)