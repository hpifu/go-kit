package htl

//func TestQueue(t *testing.T) {
//	Convey("test stack", t, func() {
//		s := NewQueue()
//		So(s, ShouldNotBeNil)
//		So(s.Len(), ShouldEqual, 0)
//		So(s.Empty(), ShouldBeTrue)
//
//		for i := 0; i < 10; i++ {
//			s.Push(i)
//			So(s.Len(), ShouldEqual, i+1)
//			So(s.Empty(), ShouldBeFalse)
//		}
//
//		for i := 0; i < 10; i++ {
//			So(s.Empty(), ShouldBeFalse)
//			So(s.Top(), ShouldEqual, i)
//			v := s.Pop()
//			So(v, ShouldNotBeNil)
//			So(v, ShouldEqual, i)
//		}
//	})
//}
