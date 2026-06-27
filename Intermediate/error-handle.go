
//  error using new method

//  Error using errorf method  ( its like a fmt.printf - here we can inject values like %d , %s )

// Error : Interface

// Go defer, panic, and recover
// defer - add delay for function execution - it will execute function at last
// multiple defere statements,
// PANIC : if our program stuck in loop of recovery and at some point we dont want to
// it to be run then we can use panic to terminate that program.
// RECOVER : sometime it is imp to run whole code even if it is fails at that time
// we recover with defer ( which basically run at the end of program ). it is used to handle panic termination. 