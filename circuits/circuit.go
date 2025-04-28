package circuits

func InputValues() ([]bool, []bool) {
	return []bool{true, true, false, false, false, true, false, true, true, false, true, false, false, true, false, true, false, true, false, true, false, true, false, true, false, false, true, false, true, false, true, true, false, false, true, false, true, false, true, false, false, false, false, false, true}, []bool{true, false, true, true, false, false, true, true, false, true, true, true, true, true, false, true, true, false, false, false, false, false, false, false, false, true, false, false, true, true, true, false, true, false, false, false, true, false, true, false, false, false, false, false, true}
}

func Circuit(x, y []bool) []bool {
	x00 := x[0]
	x01 := x[1]
	x02 := x[2]
	x03 := x[3]
	x04 := x[4]
	x05 := x[5]
	x06 := x[6]
	x07 := x[7]
	x08 := x[8]
	x09 := x[9]
	x10 := x[10]
	x11 := x[11]
	x12 := x[12]
	x13 := x[13]
	x14 := x[14]
	x15 := x[15]
	x16 := x[16]
	x17 := x[17]
	x18 := x[18]
	x19 := x[19]
	x20 := x[20]
	x21 := x[21]
	x22 := x[22]
	x23 := x[23]
	x24 := x[24]
	x25 := x[25]
	x26 := x[26]
	x27 := x[27]
	x28 := x[28]
	x29 := x[29]
	x30 := x[30]
	x31 := x[31]
	x32 := x[32]
	x33 := x[33]
	x34 := x[34]
	x35 := x[35]
	x36 := x[36]
	x37 := x[37]
	x38 := x[38]
	x39 := x[39]
	x40 := x[40]
	x41 := x[41]
	x42 := x[42]
	x43 := x[43]
	x44 := x[44]
	y00 := y[0]
	y01 := y[1]
	y02 := y[2]
	y03 := y[3]
	y04 := y[4]
	y05 := y[5]
	y06 := y[6]
	y07 := y[7]
	y08 := y[8]
	y09 := y[9]
	y10 := y[10]
	y11 := y[11]
	y12 := y[12]
	y13 := y[13]
	y14 := y[14]
	y15 := y[15]
	y16 := y[16]
	y17 := y[17]
	y18 := y[18]
	y19 := y[19]
	y20 := y[20]
	y21 := y[21]
	y22 := y[22]
	y23 := y[23]
	y24 := y[24]
	y25 := y[25]
	y26 := y[26]
	y27 := y[27]
	y28 := y[28]
	y29 := y[29]
	y30 := y[30]
	y31 := y[31]
	y32 := y[32]
	y33 := y[33]
	y34 := y[34]
	y35 := y[35]
	y36 := y[36]
	y37 := y[37]
	y38 := y[38]
	y39 := y[39]
	y40 := y[40]
	y41 := y[41]
	y42 := y[42]
	y43 := y[43]
	y44 := y[44]
	wnt := y32 && x32
	rck := y32 != x32
	skh := x15 && y15
	jqf := x15 != y15
	nhp := x02 && y02
	rvm := y02 != x02
	tcq := y24 && x24
	gww := y24 != x24
	kdd := x28 != y28
	rsv := y28 && x28
	tjq := y26 && x26
	wkb := y26 != x26
	nnn := y30 != x30
	gns := x30 && y30
	cwn := x17 && y17
	hbw := y17 != x17
	bnn := x12 && y12
	htn := y12 != x12
	tvf := y27 != x27
	kbf := x27 && y27
	gwv := x04 && y04
	jvf := y04 != x04
	rms := x31 != y31
	hjq := x31 && y31
	wmq := y33 != x33
	bfn := y33 && x33
	ncf := x41 && y41
	jhg := y41 != x41
	dbj := y43 && x43
	fhk := y43 != x43
	hns := x18 != y18
	wfm := y18 && x18
	csb := y09 && x09
	kbb := y09 != x09
	jgg := y20 && x20
	bvw := y20 != x20
	ttv := x14 && y14
	cgg := x14 != y14
	fqg := x01 && y01
	kjs := x01 != y01
	rhh := y37 && x37
	wpp := y37 != x37
	hjp := y00 && x00
	z01 := hjp != kjs
	wkq := kjs && hjp
	vdq := fqg || wkq
	hmk := vdq && rvm
	wvt := hmk || nhp
	z02 := rvm != vdq
	z00 := x00 != y00
	gvj := y10 && x10
	rmb := x10 != y10
	nss := x34 != y34
	bnw := y34 && x34
	hhw := x06 && y06
	jtg := x06 != y06
	dmp := x13 != y13
	rbg := x13 && y13
	djt := x35 != y35
	qtv := x35 && y35
	jdc := y36 && x36
	hbh := y36 != x36
	rjs := y08 != x08
	gdd := x08 && y08
	qpc := x07 != y07
	jbw := x07 && y07
	cng := y42 != x42
	hnh := x42 && y42
	nns := x40 != y40
	bdn := x40 && y40
	ppk := y39 && x39
	vqf := y39 != x39
	bsw := y29 && x29
	hvc := y29 != x29
	hpn := y05 && x05
	fds := x05 != y05
	gfk := x03 && y03
	sbt := y03 != x03
	dpv := wvt && sbt
	jth := gfk || dpv
	whn := jth && jvf
	whf := gwv || whn
	z05 := whf != fds
	vvs := fds && whf
	cfn := vvs || hpn
	z06 := jtg != cfn
	qhk := jtg && cfn
	qmb := hhw || qhk
	z07 := qmb != qpc
	trv := qpc && qmb
	tmg := jbw || trv
	z08 := rjs != tmg
	vth := tmg && rjs
	wpb := vth || gdd
	z09 := wpb != kbb
	jmp := kbb && wpb
	kmr := jmp || csb
	fws := rmb && kmr
	qqw := fws || gvj
	z10 := kmr != rmb
	z04 := jvf != jth
	z03 := wvt != sbt
	cqg := y25 && x25
	prp := y25 != x25
	fnc := y23 && x23
	mcp := x23 != y23
	cmp := x19 != y19
	z19 := y19 && x19
	gmf := y22 && x22
	wwp := y22 != x22
	gkc := y11 != x11
	z11 := gkc && qqw
	wpd := qqw != gkc
	dpf := y11 && x11
	dtq := wpd || dpf
	z12 := htn != dtq
	gvh := htn && dtq
	jkm := bnn || gvh
	qpw := dmp && jkm
	rhf := qpw || rbg
	z14 := cgg != rhf
	smd := rhf && cgg
	rkt := smd || ttv
	kjk := rkt && skh
	kbq := jqf || kjk
	z15 := skh != rkt
	z13 := jkm != dmp
	jsp := x44 && y44
	jsg := x44 != y44
	rhd := x16 && y16
	rvn := y16 != x16
	vtm := kbq && rvn
	wrc := rhd || vtm
	jks := hbw && wrc
	qvq := jks || cwn
	mts := hns && qvq
	wfc := wfm || mts
	pbb := cmp && wfc
	mdd := wfc != cmp
	hvn := pbb || mdd
	cmm := bvw && hvn
	qpm := jgg || cmm
	z20 := bvw != hvn
	z18 := qvq != hns
	z17 := wrc != hbw
	z16 := rvn != kbq
	sqj := x38 != y38
	jhv := y38 && x38
	nrg := y21 && x21
	csw := y21 != x21
	mvs := csw && qpm
	bvr := nrg || mvs
	z22 := wwp != bvr
	cgv := wwp && bvr
	mkf := cgv || gmf
	mmk := mcp && mkf
	jmf := fnc || mmk
	rjb := jmf && gww
	pfc := tcq || rjb
	z25 := pfc != prp
	rkh := prp && pfc
	jkr := rkh || cqg
	hhj := wkb && jkr
	cmb := tjq || hhj
	mqr := cmb && tvf
	msf := kbf || mqr
	z28 := msf != kdd
	dpr := kdd && msf
	tsk := rsv || dpr
	z29 := tsk != hvc
	bfp := hvc && tsk
	pwp := bsw || bfp
	z30 := pwp != nnn
	dwg := pwp && nnn
	sgf := gns || dwg
	mrg := sgf && rms
	ftq := mrg || hjq
	z32 := rck != ftq
	hfc := rck && ftq
	dts := wnt || hfc
	sbw := dts && wmq
	jmv := sbw || bfn
	z34 := jmv != nss
	vtd := nss && jmv
	khf := vtd || bnw
	z35 := djt != khf
	khs := khf && djt
	rfq := qtv || khs
	pwb := rfq && hbh
	smt := pwb || jdc
	jgw := wpp && smt
	z37 := jgw || rhh
	wts := smt != wpp
	z38 := sqj != wts
	qdn := sqj && wts
	fkq := qdn || jhv
	hwb := vqf && fkq
	jbd := ppk || hwb
	z40 := nns != jbd
	nnq := jbd && nns
	gfd := nnq || bdn
	z41 := jhg != gfd
	bbr := jhg && gfd
	mwt := bbr || ncf
	hbm := mwt && cng
	gsk := hbm || hnh
	tnc := gsk && fhk
	hks := tnc || dbj
	nrv := jsg && hks
	z45 := nrv || jsp
	z44 := jsg != hks
	z43 := gsk != fhk
	z42 := cng != mwt
	z39 := vqf != fkq
	z36 := rfq != hbh
	z33 := dts != wmq
	z31 := rms != sgf
	z27 := tvf != cmb
	z26 := wkb != jkr
	z24 := gww != jmf
	z23 := mcp != mkf
	z21 := qpm != csw
	return []bool{z00, z01, z02, z03, z04, z05, z06, z07, z08, z09, z10, z11, z12, z13, z14, z15, z16, z17, z18, z19, z20, z21, z22, z23, z24, z25, z26, z27, z28, z29, z30, z31, z32, z33, z34, z35, z36, z37, z38, z39, z40, z41, z42, z43, z44, z45}
}
