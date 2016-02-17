// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dist

import (
	"math"
	"testing"
)

func TestLaplaceProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0,
			prob:    0.5,
			cumProb: 0.5,
			logProb: math.Log(0.5),
		},
		{
			loc:     -1,
			prob:    1 / (2 * math.E),
			cumProb: 0.1839397205857211607977618850807304337229055655158839172539184008487307478724499016785736371729598219,
			logProb: math.Log(1 / (2 * math.E)),
		},
		{
			loc:     1,
			prob:    1 / (2 * math.E),
			cumProb: 0.8160602794142788392022381149192695662770944344841160827460815991512692521275500983214263628270401781,
			logProb: math.Log(1 / (2 * math.E)),
		},
		{
			loc:     -7,
			prob:    1 / (2 * math.Pow(math.E, 7)),
			cumProb: 0.0004559409827772581040015680422046413132368622637180269204080667109447399446551532646631395032324502210,
			logProb: math.Log(1 / (2 * math.Pow(math.E, 7))),
		},
		{
			loc:     7,
			prob:    1 / (2 * math.Pow(math.E, 7)),
			cumProb: 0.9995440590172227418959984319577953586867631377362819730795919332890552600553448467353368604967675498,
			logProb: math.Log(1 / (2 * math.Pow(math.E, 7))),
		},
		{
			loc:     -20,
			prob:    math.Exp(-20.69314718055994530941723212145817656807550013436025525412068000949339362196969471560586332699641869),
			cumProb: 1.030576811219278913982970190077910488187903637799551846486122330814582011892279676639955463952790684 * 1e-9,
			logProb: -20.69314718055994530941723212145817656807550013436025525412068000949339362196969471560586332699641869,
		},
		{
			loc:     20,
			prob:    math.Exp(-20.69314718055994530941723212145817656807550013436025525412068000949339362196969471560586332699641869),
			cumProb: 0.999999998969423188780721086017029809922089511812096362200448153513877669185417988107720323360044536,
			logProb: -20.69314718055994530941723212145817656807550013436025525412068000949339362196969471560586332699641869,
		},
	}
	testDistributionProbs(t, Laplace{Mu: 0, Scale: 1}, "Laplace", pts)
}
