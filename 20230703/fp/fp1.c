#include <stdio.h>
#include <stdlib.h>

int compare(const void *pa, const void *pb){
  int a = *(const int *)pa;
  int b = *(const int *)pb;
  return a - b;
}

int compareMagnitude(const void *pa, const void *pb){
  int a = (unsigned int)*(const int *)pa;
  int b = (unsigned int)*(const int *)pb;
  return abs(a) - abs(b);
}
int main(int argc, char **argv){
  int arr[] = {1, -4, 4, 39, -5, 9};
  qsort(arr, sizeof(arr)/sizeof(arr[0]), sizeof(arr[0]), compare);
  for (int i=0; i < sizeof(arr)/sizeof(arr[0]); i++)
    printf("%d ", arr[i]);
  printf("\nBy magnitudes...\n");
    qsort(arr, sizeof(arr)/sizeof(arr[0]), sizeof(arr[0]), compareMagnitude);
  for (int i=0; i < sizeof(arr)/sizeof(arr[0]); i++)
    printf("%d ", arr[i]);
  return 0;
}
