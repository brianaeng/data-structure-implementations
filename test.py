#List comprehension: http://treyhunner.com/2015/12/python-list-comprehensions-now-in-color/

#Removing things from a list (aka creating a new list for items I want)
def find_one(arr):
    final_array = [x for x in arr if x == 1]
    print final_array
    return final_array

find_one([2,3,1,1,3,1,5,2,1,5,1,7,8,1,1,1,8])

#Replacing things in a list (aka creating a new list for items I want)
def new_values(arr):
    final_array = [x*2 for x in arr]
    print final_array

    final_array2 = [x*2 for x in arr if x==1]
    print final_array2

new_values([2,1,1,2,2,1])

#Looping through an arraylist with value
def lets_loop_array(arr):
    for x in arr:
        print x

lets_loop_array([1,2,3,4,5,6])

#Looping through an arraylist with index and value
def lets_loop_array2(arr):
    for index, val in enumerate(arr):
        print index, val

lets_loop_array2([1,2,3,4])

#Looping through a hashmap with key
def lets_loop_map(hashy):
    for key in hashy:
        print key

lets_loop_map({1:"one", 2:"two", 3:"three"})

#Looping through a hashmap with key and value
def lets_loop_map2(hashy):
    for key, value in hashy.iteritems():
        print key, value

lets_loop_map2({1:"one", 2:"two", 3:"three"})