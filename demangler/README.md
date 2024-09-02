# Demangler

This all assumes g++ is the compiler in use for the g++-specific vnendor 'cxxabi.h' header to find meaningful ABI symbol names.

Notably, I could not have written this without AI assistance. There were just too many ugly pitfalls in C++ with reference semantics and I ultimately gave up on a solution with universal forwarding and opted for a parametrized macro (again, macros are something AI brings here).

This was something I began thinking about to answer some long standing questions about expression state when universal forwarding.
