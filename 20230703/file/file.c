#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main(){
  FILE *fp = fopen("doc", "r");
  char ch = fgetc(fp);
  bool start = true;
  int i = 1;
  while (ch != EOF) {
    if (start == true) {
      printf("%d # ", i);
      i++;
      start = false;
    }
    if (ch == '\n'){
      start = true;
    }
    printf("%c", ch);
    ch = fgetc(fp);
  }
  fclose(fp);
  return 0;
}
