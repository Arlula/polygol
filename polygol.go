package polygol

type Geom [][][][]float64

type Polygol struct{}

func New() *Polygol {
	return &Polygol{}
}

func (p *Polygol) Union(geom Geom, moreGeoms ...Geom) (Geom, error) {
	return newOperation(opUnion).run(geom, moreGeoms...)
}

func (p *Polygol) Intersection(geom Geom, moreGeoms ...Geom) (Geom, error) {
	return newOperation(opIntersection).run(geom, moreGeoms...)
}

func (p *Polygol) Difference(geom Geom, moreGeoms ...Geom) (Geom, error) {
	return newOperation(opDifference).run(geom, moreGeoms...)
}

func (p *Polygol) XOR(geom Geom, moreGeoms ...Geom) (Geom, error) {
	return newOperation(opXor).run(geom, moreGeoms...)
}

func Union(geom Geom, moreGeoms ...Geom) (Geom, error) {
	return New().Union(geom, moreGeoms...)
}

func Intersection(geom Geom, moreGeoms ...Geom) (Geom, error) {
	return New().Intersection(geom, moreGeoms...)
}

func Difference(geom Geom, moreGeoms ...Geom) (Geom, error) {
	return New().Difference(geom, moreGeoms...)
}

func XOR(geom Geom, moreGeoms ...Geom) (Geom, error) {
	return New().XOR(geom, moreGeoms...)
}
