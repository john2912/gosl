// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"math"

	"github.com/cpmech/gosl/utl"
)

// newERKhighOrder returns the coefficients of the explicit Runge-Kutta method (high order)
func newERKhighOrder(kind string) rkmethod {

	// new dataset
	o := new(ExplicitRK)

	// set coefficients
	switch kind {

	case "fehlberg7": // Fehlberg 7(8) ⇒ q = 7
		o.Embedded = true
		o.A = [][]float64{
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{2.0 / 27.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{1.0 / 36.0, 1.0 / 12.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{1.0 / 24.0, 0.0, 1.0 / 8.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{5.0 / 12.0, 0.0, -25.0 / 16.0, 25.0 / 16.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{1.0 / 20.0, 0.0, 0.0, 1.0 / 4.0, 1.0 / 5.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{-25.0 / 108.0, 0.0, 0.0, 125.0 / 108.0, -65.0 / 27.0, 125.0 / 54.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{31.0 / 300.0, 0.0, 0.0, 0.0, 61.0 / 225.0, -2.0 / 9.0, 13.0 / 900.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{2.0, 0.0, 0.0, -53.0 / 6.0, 704.0 / 45.0, -107.0 / 9.0, 67.0 / 90.0, 3.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{-91.0 / 108.0, 0.0, 0.0, 23.0 / 108.0, -976.0 / 135.0, 311.0 / 54.0, -19.0 / 60.0, 17.0 / 6.0, -1.0 / 12.0, 0.0, 0.0, 0.0, 0.0},
			{2383.0 / 4100.0, 0.0, 0.0, -341.0 / 164.0, 4496.0 / 1025.0, -301.0 / 82.0, 2133.0 / 4100.0, 45.0 / 82.0, 45.0 / 164.0, 18.0 / 41.0, 0.0, 0.0, 0.0},
			{3.0 / 205.0, 0.0, 0.0, 0.0, 0.0, -6.0 / 41.0, -3.0 / 205.0, -3.0 / 41.0, 3.0 / 41.0, 6.0 / 41.0, 0.0, 0.0, 0.0},
			{-1777.0 / 4100.0, 0.0, 0.0, -341.0 / 164.0, 4496.0 / 1025.0, -289.0 / 82.0, 2193.0 / 4100.0, 51.0 / 82.0, 33.0 / 164.0, 12.0 / 41.0, 0.0, 1.0, 0.0},
		}
		o.B = []float64{41.0 / 840.0, 0.0, 0.0, 0.0, 0.0, 34.0 / 105.0, 9.0 / 35.0, 9.0 / 35.0, 9.0 / 280.0, 9.0 / 280.0, 41.0 / 840.0, 0.0, 0.0}
		o.Be = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 34.0 / 105.0, 9.0 / 35.0, 9.0 / 35.0, 9.0 / 280.0, 9.0 / 280.0, 0.0, 41.0 / 840.0, 41.0 / 840.0}
		o.C = []float64{0.0, 2.0 / 27.0, 1.0 / 9.0, 1.0 / 6.0, 5.0 / 12.0, 1.0 / 2.0, 5.0 / 6.0, 1.0 / 6.0, 2.0 / 3.0, 1.0 / 3.0, 1.0, 0.0, 1.0}
		o.E = []float64{41.0 / 840.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 41.0 / 840.0, -41.0 / 840.0, -41.0 / 840.0}
		o.P = 7
		o.Q = 8

	case "dopri8":
		o.Embedded = true
		o.P = 8
		o.Q = 7

		SQ6 := math.Sqrt(6.0)
		sm6 := 6.0 - SQ6
		sp6 := 6.0 + SQ6
		o.C = []float64{
			0.0,               // c1
			2.0 * sm6 / 135.0, // c2
			sm6 / 45.0,        // c3
			sm6 / 30.0,        // c4
			sp6 / 30.0,        // c5
			1.0 / 3.0,         // c6
			1.0 / 4.0,         // c7
			4.0 / 13.0,        // c8
			127.0 / 195.0,     // c9
			3.0 / 5.0,         // c10
			6.0 / 7.0,         // c11
			1.0,               // c12
			/* TODO: dense output
			0.0,               // c13
			0.1,               // c14
			0.2,               // c15
			7.0 / 9.0,         // c16
			*/
		}

		o.B = []float64{
			5.42937341165687622380535766363e-2,  // b1
			0.0000000000000000000000000000000,   // b2
			0.0000000000000000000000000000000,   // b3
			0.0000000000000000000000000000000,   // b4
			0.0000000000000000000000000000000,   // b5
			4.45031289275240888144113950566e0,   // b6
			1.89151789931450038304281599044e0,   // b7
			-5.8012039600105847814672114227e0,   // b8
			3.1116436695781989440891606237e-1,   // b9
			-1.52160949662516078556178806805e-1, // b10
			2.01365400804030348374776537501e-1,  // b11
			4.47106157277725905176885569043e-2,  // b12
			/* TODO: dense output
			0.00000000000000000000000000000000,   // b13
			0.244094488188976377952755905512e+00, // b14 := bhh1
			0.733846688281611857341361741547e+00, // b15 := bhh2
			0.220588235294117647058823529412e-01, // b16 := bhh3
			*/
		}

		o.A = utl.Alloc(12, 12)
		o.A[1][0] = 5.26001519587677318785587544488e-2  //  2,1
		o.A[2][0] = 1.97250569845378994544595329183e-2  //  3,1
		o.A[2][1] = 5.91751709536136983633785987549e-2  //  3,2
		o.A[3][0] = 2.95875854768068491816892993775e-2  //  4,1
		o.A[3][2] = 8.87627564304205475450678981324e-2  //  4,3
		o.A[4][0] = 2.41365134159266685502369798665e-1  //  5,1
		o.A[4][2] = -8.84549479328286085344864962717e-1 //  5,3
		o.A[4][3] = 9.24834003261792003115737966543e-1  //  5,4

		o.A[5][0] = 3.7037037037037037037037037037e-2  //  6,1
		o.A[5][3] = 1.70828608729473871279604482173e-1 //  6,4
		o.A[5][4] = 1.25467687566822425016691814123e-1 //  6,5

		o.A[6][0] = 3.71093750000000000000000000000e-2  //  7,1
		o.A[6][3] = 1.70252211019544039314978060272e-1  //  7,4
		o.A[6][4] = 6.02165389804559606850219397283e-2  //  7,5
		o.A[6][5] = -1.75781250000000000000000000000e-2 //  7,6

		o.A[7][0] = 3.70920001185047927108779319836e-2  //  8,1
		o.A[7][3] = 1.70383925712239993810214054705e-1  //  8,4
		o.A[7][4] = 1.07262030446373284651809199168e-1  //  8,5
		o.A[7][5] = -1.53194377486244017527936158236e-2 //  8,6
		o.A[7][6] = 8.27378916381402288758473766002e-3  //  8,7

		o.A[8][0] = 6.24110958716075717114429577812e-1  //  9,1
		o.A[8][3] = -3.36089262944694129406857109825e0  //  9,4
		o.A[8][4] = -8.68219346841726006818189891453e-1 //  9,5
		o.A[8][5] = 2.75920996994467083049415600797e1   //  9,6
		o.A[8][6] = 2.01540675504778934086186788979e1   //  9,7
		o.A[8][7] = -4.34898841810699588477366255144e1  //  9,8

		o.A[9][0] = 4.77662536438264365890433908527e-1  // 10,1
		o.A[9][3] = -2.48811461997166764192642586468e0  // 10,4
		o.A[9][4] = -5.90290826836842996371446475743e-1 // 10,5
		o.A[9][5] = 2.12300514481811942347288949897e1   // 10,6
		o.A[9][6] = 1.52792336328824235832596922938e1   // 10,7
		o.A[9][7] = -3.32882109689848629194453265587e1  // 10,8
		o.A[9][8] = -2.03312017085086261358222928593e-2 // 10,9

		o.A[10][0] = -9.3714243008598732571704021658e-1 // 11,1
		o.A[10][3] = 5.18637242884406370830023853209e0  // 11,4
		o.A[10][4] = 1.09143734899672957818500254654e0  // 11,5
		o.A[10][5] = -8.14978701074692612513997267357e0 // 11,6
		o.A[10][6] = -1.85200656599969598641566180701e1 // 11,7
		o.A[10][7] = 2.27394870993505042818970056734e1  // 11,8
		o.A[10][8] = 2.49360555267965238987089396762e0  // 11,9
		o.A[10][9] = -3.0467644718982195003823669022e0  // 11,10

		o.A[11][0] = 2.27331014751653820792359768449e0   // 12,1
		o.A[11][3] = -1.05344954667372501984066689879e1  // 12,4
		o.A[11][4] = -2.00087205822486249909675718444e0  // 12,5
		o.A[11][5] = -1.79589318631187989172765950534e1  // 12,6
		o.A[11][6] = 2.79488845294199600508499808837e1   // 12,7
		o.A[11][7] = -2.85899827713502369474065508674e0  // 12,8
		o.A[11][8] = -8.87285693353062954433549289258e0  // 12,9
		o.A[11][9] = 1.23605671757943030647266201528e1   // 12,10
		o.A[11][10] = 6.43392746015763530355970484046e-1 // 12,11

		/* TODO: dense output
		o.A[13][0] = 5.61675022830479523392909219681e-2  // 14,1
		o.A[13][6] = 2.53500210216624811088794765333e-1  // 14,7
		o.A[13][7] = -2.46239037470802489917441475441e-1 // 14,8
		o.A[13][8] = -1.24191423263816360469010140626e-1 // 14,9
		o.A[13][9] = 1.5329179827876569731206322685e-1   // 14,10
		o.A[13][10] = 8.20105229563468988491666602057e-3 // 14,11
		o.A[13][11] = 7.56789766054569976138603589584e-3 // 14,12
		o.A[13][12] = -8.2980000000000000000000000000e-3 // 14,13

		o.A[14][0] = 3.18346481635021405060768473261e-2   // 15,1
		o.A[14][5] = 2.83009096723667755288322961402e-2   // 15,6
		o.A[14][6] = 5.35419883074385676223797384372e-2   // 15,7
		o.A[14][7] = -5.49237485713909884646569340306e-2  // 15,8
		o.A[14][10] = -1.08347328697249322858509316994e-4 // 15,11
		o.A[14][11] = 3.82571090835658412954920192323e-4  // 15,12
		o.A[14][12] = -3.40465008687404560802977114492e-4 // 15,13
		o.A[14][13] = 1.41312443674632500278074618366e-1  // 15,14

		o.A[15][0] = -4.28896301583791923408573538692e-1  // 16,1
		o.A[15][5] = -4.69762141536116384314449447206e0   // 16,6
		o.A[15][6] = 7.68342119606259904184240953878e0    // 16,7
		o.A[15][7] = 4.06898981839711007970213554331e0    // 16,8
		o.A[15][8] = 3.56727187455281109270669543021e-1   // 16,9
		o.A[15][12] = -1.39902416515901462129418009734e-3 // 16,13
		o.A[15][13] = 2.9475147891527723389556272149e0    // 16,14
		o.A[15][14] = -9.15095847217987001081870187138e0  // 16,15
		*/

		o.E = []float64{
			0.1312004499419488073250102996e-01,  // er1
			0.00000000000000000000000000000000,  // er2
			0.00000000000000000000000000000000,  // er3
			0.00000000000000000000000000000000,  // er4
			0.00000000000000000000000000000000,  // er5
			-0.1225156446376204440720569753e+01, // er6
			-0.4957589496572501915214079952e+00, // er7
			0.1664377182454986536961530415e+01,  // er8
			-0.3503288487499736816886487290e+00, // er9
			0.3341791187130174790297318841e+00,  // er10
			0.8192320648511571246570742613e-01,  // er11
			-0.2235530786388629525884427845e-01, // er12
		}

	default:
		return nil
	}

	// set number of stages
	o.Nstg = len(o.A)
	return o
}
