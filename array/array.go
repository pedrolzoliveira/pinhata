package array

type Array[T any] []T

func (arr *Array[T]) ForEach(function func(T)) {
	for _, item := range *arr {
		function(item)
	}
}

func (arr *Array[T]) Filter(function func(T) bool) Array[T] {
	newArr := make(Array[T], 0)
	for _, item := range *arr {
		if function(item) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

func (arr *Array[T]) Add(item T) {
	*arr = append(*arr, item)
}

func (arr *Array[T]) Has(function func(T) bool) bool {
	for _, item := range *arr {
		if function(item) {
			return true
		}
	}
	return false
}

func (arr *Array[T]) Find(function func(T) bool) (bool, T) {
	for _, item := range *arr {
		if function(item) {
			return true, item
		}
	}

	var zeroValue T
	return false, zeroValue
}
