#include <stdlib.h>
#include <stdio.h>
//#include <string>
int compare (const void *pa, const void *pb){
  int a = *(const int*)pa;
  int b = *(const int*)pb;
  return a == b
    ? 0 : a < b
      ? 1 : -1;
}
int partition (int *nums, int len, int l, int r) {
  int i = l - 1, k = 0, p = nums[r];
  for (int x = l; x < r; x++) {
    if compare(arr + x, arr + r) {
	//TODO
    }
  }
}
void qsort1(int *nums, int len, int (*cmp) (void *, void *)) {
  //TODO
}
//void printList(int
int main(int argc, char **argv){
  printList(argv, argc);
  qsort1(argv, argc, compare);
  printList(argv, argc);
}

int main2(){
  printf("Input ");
  char *str_num = malloc(6);
  gets(str_num);
  int m = atoi(str_num);
  printf("\t%d\n", 2*m);
  return 0;
}
