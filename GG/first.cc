#include <iostream>
#include <exception>
#include <signal.h>
#include <stdlib.h>

void reverse(std::string &line)
{
  // // Approach #1
  // auto lt = line.begin();
  // auto rt = line.end();
  // reverse(lt, rt);

  // // Approach #2
  // for (int i = 0; i < line.length() / 2; i++)
  // {
  //   std::swap(line[i], line[line.length() - 1 - i]);
  // }
  char *reversed = static_cast<char *>(malloc(sizeof(char) * (line.length() + 2)));
  memset(reversed, 0, line.length() + 2);
  int i = line.length();
  for (auto c : line)
  {
    i--;
    if (i < 0)
      break;
    reversed[i] = c;
  }
  for (int i = 0; i < line.length(); i++)
  {
    line[i] = reversed[i];
  }
  free(reversed);
}
void wmain()
{
  std::string inputLine;
  while (1)
  {
    std::cout << "Enter a string: ";
    getline(std::cin, inputLine);
    reverse(inputLine);
    std::cout << inputLine << std::endl;
  }
}

void handle_exit(int signal)
{
  throw "Danger Will Robinson";
}

int main()
{
  double x;
  double sum = 0.0;
  struct sigaction s_action;
  s_action.sa_handler = handle_exit;
  s_action.sa_flags = 0;
  sigemptyset(&s_action.sa_mask);
  sigaction(SIGINT, &s_action, NULL);
  sigaction(SIGSTOP, &s_action, NULL);
  std::string words = "";
  // SIGI NT
  bool firstRun = true;
  std::string line;
  std::string x_str;

  wmain();

  return 0;
}
