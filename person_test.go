package main

import "testing"

func TestTwiceOlderExists(t *testing.T) {

	t.Run("false response if not number exactly twice as big", func(t *testing.T) {
		persons := []Person{
			{1},
			{4},
			{10},
		}

		expected := false
		ret := TwiceOlderExists(persons)
		if expected != ret {
			t.Errorf("expected %t but got %t", expected, ret)
		}
	})

	t.Run("false response if not number extactly twice as big", func(t *testing.T) {
		persons := []Person{
			{1},
			{4},
			{8},
		}
		expected := true
		ret := TwiceOlderExists(persons)
		if expected != ret {
			t.Errorf("expected %t but got %t", expected, ret)
		}
	})

	t.Run("false response if not number extactly twice as big", func(t *testing.T) {
		// pass emtpy slice
		persons := []Person{}
		expected := false
		ret := TwiceOlderExists(persons)
		if expected != ret {
			t.Errorf("expected %t but got %t", expected, ret)
		}
	})

}

func TestMinTwiceOlderExists(t *testing.T) {

	t.Run("false response if there is not number that is at least exactly twice as big", func(t *testing.T) {
		persons := []Person{
			{1},
			{4},
			{5},
		}

		expected := false
		ret := MinTwiceOlderExists(persons)
		if expected != ret {
			t.Errorf("expected %t but got %t", expected, ret)
		}
	})

	t.Run("false response if there is not number that is at least exactly twice as big", func(t *testing.T) {
		persons := []Person{
			{1},
			{4},
			{10},
		}
		expected := true
		ret := MinTwiceOlderExists(persons)
		if expected != ret {
			t.Errorf("expected %t but got %t", expected, ret)
		}
	})

	t.Run("false response if there is not number that is at least exactly twice as big", func(t *testing.T) {
		// pass emtpy slice
		persons := []Person{}
		expected := false
		ret := MinTwiceOlderExists(persons)
		if expected != ret {
			t.Errorf("expected %t but got %t", expected, ret)
		}
	})

}
