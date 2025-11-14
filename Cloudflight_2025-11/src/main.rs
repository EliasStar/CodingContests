mod level1;
mod level2;
mod level3;
mod level4;

fn main() {
    level1::run().expect("level1 not working");
    level2::run().expect("level2 not working");
    level3::run().expect("level3 not working");
    level4::run().expect("level4 not working");
}
