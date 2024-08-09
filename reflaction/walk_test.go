package reflaction

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	name := "Arjesh"
	age := 26
	city := "Surat"

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{name},
			[]string{name},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{name, city},
			[]string{name, city},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{name, age},
			[]string{name},
		},
		{
			"nested fields",
			Person{
				name,
				Profile{age, city},
			},
			[]string{name, city},
		},
		{
			"pointers to things",
			&Person{
				name,
				Profile{age, city},
			},
			[]string{name, city},
		},
		{
			"slices",
			[]Profile{
				{age, city},
				{27, "Ahmedabad"},
			},
			[]string{city, "Ahmedabad"},
		},
		{
			"arrays",
			[2]Profile{
				{age, city},
				{27, "Ahmedabad"},
			},
			[]string{city, "Ahmedabad"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{age, city}
			aChannel <- Profile{27, "Ahmedabad"}
			close(aChannel)
		}()

		var got []string
		want := []string{city, "Ahmedabad"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{age, city}, Profile{27, "Ahmedabad"}
		}

		var got []string
		want := []string{city, "Ahmedabad"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
