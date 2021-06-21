package metrics

type Bucket struct {
	Points []float64
	Count  int64
	next   *Bucket
}

func (b *Bucket) Append(val float64) {
	b.Points = append(b.Points, val)
	b.Count++

}

func (b *Bucket) Add(offset int, val float64) {
	b.Points[offset] += val
	b.Count++

}

func (b *Bucket) Reset() {
	b.Points = b.Points[:0]
	b.Count = 0

}

func (b *Bucket) Next() *Bucket {
	return b.next

}

type Window struct {
	window []Bucket
	size   int
}

type WindowOpts struct {
	Size int
}

func NewWindow(opts WindowOpts) *Window {
	buckets := make([]Bucket, opts.Size)
	for offset := range buckets {
		buckets[offset] = Bucket{Points: make([]float64, 0)}
		nextOffset := offset + 1
		if nextOffset == opts.Size {
			nextOffset = 0
		}
		buckets[offset].next = &buckets[nextOffset]
	}
	return &Window{window: buckets, size: opts.Size}
}

func (w *Window) ResetWindow()  {
	for offset := range w.window {
		w.ResetBucket(offset)
	}

}

func (w *Window) ResetBucket(offset int)  {
	w.window[offset].Reset()

}

func (w *Window) ResetBuckets(offsets []int)  {
	for _, offset := range offsets {
		w.ResetBucket(offset)
	}

}

func (w *Window) Append(offset int, val float64)  {
	w.window[offset].Append(val)

}

func (w *Window) Add(offset int, val float64)  {
	if w.window[offset].Count == 0 {
		w.window[offset].Append(val)
		return
	}
	w.window[offset].Add(0, val)
}

func (w *Window) Bucket(offset int) Bucket {
	return w.window[offset]
}

func (w *Window) Size() int {
	return w.size
}

func (w *Window) Iterator(offset int, count int) Iterator {
	return Iterator{
		count: count,
		cur: &w.window[offset],
	}
}