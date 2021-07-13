#include<iostream>

using namespace std; 

int lowerb(int arr[] , int start, int end, int x) {
    while (start < end) {
        int mid = (start + end) >> 1;
        if (x <= arr[mid]){
            end = mid;
        } else {
            start = mid + 1;
        }
    }
    return start;
}

int upperb(int arr[] , int start, int end, int x) {
    while (start < end) {
        int mid = (start + end) >> 1;
        if (x >= arr[mid]){
            start = mid + 1;
        } else {
            end = mid;
        }
    }
    return start;
}

int main() {
    int arr[8] = {1, 3, 3, 3, 4, 5, 6,};
    cout << lowerb(arr,0,8,2) << endl; // return 1
    cout << upperb(arr,0,8,2) << endl; // return 1
    cout << lowerb(arr,0,8,3) << endl; // return 1
    cout << upperb(arr,0,8,3) << endl; // return 4

    
}
