#include <iostream>

int *returnInterger()
{
  int x = 100;
  return &x;
}

int main()
{
  std::cout << "This is C Pointer Test Appliction\n";
  int *x = returnInterger();
  std::cout << *x;
}
