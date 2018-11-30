package pso

import (
	"testing"

	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
)

func Test_NewParticle(t *testing.T) {
	p := pp.PricingProblem{}
	pr := *p.MakeProblem(2, false)
	sw := NewSwarm(2, 1, &pr)
	if len(sw.Particles[0].prices) != 2 {
		t.Errorf("particle prices design length not as expected : %v", len(sw.Particles[0].prices))
	}
	if sw.Particles[0].currentRevenue == 0.0 {
		t.Errorf("particle current revenue not 0 : %v", sw.Particles[0].currentRevenue)
	}
	if len(sw.Particles[0].velocity) != 2 {
		t.Errorf("particle velocity length not 2 : %v", len(sw.Particles[0].velocity))
	}
	if len(sw.Particles[0].bestPrices) != 2 {
		t.Errorf("particle best price length not 2 : %v", len(sw.Particles[0].bestPrices))
	}
	if sw.Particles[0].bestRevenue == 0.0 {
		t.Errorf("particle length not calculated : %v", sw.Particles[0].bestRevenue)
	}
}

func Test_calculateVelocity(t *testing.T) {
	p1 := []float64{0.1, 0.5, 1, 3}
	p2 := []float64{0.2, 1, 2, 6}
	v := initialVelocity(p1, p2)
	newV := calculateVelocity(v, p1, p1, []float64{0.5, 3.2, 2.1, 0.2})
	if newV[0] == v[0] {
		t.Errorf("new 0 calculated velocity did not change : %v", newV[0])
	}
	if newV[1] == v[1] {
		t.Errorf("new 1 calculated velocity did not change : %v", newV[1])
	}
}

func Test_initialVelocity(t *testing.T) {
	p1 := []float64{0.25, 0.5}
	p2 := []float64{0.5, 1.0}
	v := initialVelocity(p1, p2)
	if v[0] != 0.125 {
		t.Errorf("expected velocity[0] %v, actual %v", 0.125, v[0])
	}
	if v[1] != 0.25 {
		t.Errorf("expected velocity[0] %v, actual %v", 0.25, v[1])
	}
}

func Test_randomDesign(t *testing.T) {
	p := pp.PricingProblem{}
	pr := *p.MakeProblem(2, false)
	prices := randomPrices(2, &pr)
	if !pr.IsValid(prices) {
		t.Errorf("invalid prices found : %v", prices)
	}
}

func Test_updatePosition(t *testing.T) {
	p1 := []float64{0.25, 0.5}
	p2 := []float64{0.5, 1.0}
	v := initialVelocity(p1, p2)
	p := pp.PricingProblem{}
	pr := *p.MakeProblem(2, false)

	np := updatePosition(p1, v, 3, &pr)
	if len(np) != len(p1) {
		t.Errorf("incorrect position length : %v", len(np))
	}
	if np[0] == p1[0] {
		t.Errorf("new design 0 did not update : %v", np[0])
	}
	if np[1] == p1[1] {
		t.Errorf("new design 1 did not update : %v", np[1])
	}
}
