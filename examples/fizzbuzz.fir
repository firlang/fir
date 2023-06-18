io := require("stdio");

fn main() {
    i := 0;
    for i <= 20 {
        if i % 3 == 0 && i % 5 == 0 {
            io.println("FizzBuzz");
        } else if i % 3 == 0 {
            io.println("Fizz");
        } else if i % 5 == 0 {
            io.println("Buzz");
        } else {
            io.println("{}", i)
        }
    }
}