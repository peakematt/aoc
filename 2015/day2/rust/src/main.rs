use std::cmp::Ordering;
use std::fs;

struct Present {
    length: i32,
    width: i32,
    height: i32,
}

#[derive(Eq, Clone, Copy)]
struct Side {
    length: i32,
    width: i32,
}

impl Side {
    fn area(&self) -> i32 {
        self.length * self.width
    }

    fn perimeter(&self) -> i32 {
        2 * self.length + 2 * self.width
    }
}

impl Ord for Side {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.area().cmp(&other.area())
    }
}

impl PartialOrd for Side {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl PartialEq for Side {
    fn eq(&self, other: &Self) -> bool {
        self.length == other.length && self.width == other.width
    }
}

impl Present {
    fn sides(&self) -> Vec<Side> {
        vec![
            Side {
                length: self.length,
                width: self.width,
            },
            Side {
                length: self.width,
                width: self.height,
            },
            Side {
                length: self.length,
                width: self.height,
            },
        ]
    }

    fn surface_area(&self) -> i32 {
        // 2*(self.length * self.width) + 2*(self.width * self.height) + 2*(self.length * self.height)
        self.sides().iter().map(|s: &Side| 2 * s.area()).sum()
    }

    fn smallest_surface_area_side(&self) -> Side {
        let mut smallest_side = self.sides();
        smallest_side.sort();
        smallest_side[0]
    }

    fn surface_area_and_extra(&self) -> i32 {
        self.surface_area() + self.smallest_surface_area_side().area()
    }

    fn volume(&self) -> i32 {
        self.width * self.length * self.height
    }

    fn ribbon_needed(&self) -> i32 {
        self.smallest_surface_area_side().perimeter() + self.volume()
    }
}

fn main() {
    let input = fs::read_to_string("src/presents.txt").expect("Failed to read input file");
    let presents: Vec<Present> = input
        .lines()
        .map(|line| {
            let dimensions: Vec<i32> = line.split("x").map(|x| x.parse().unwrap()).collect();
            Present {
                length: dimensions[0],
                width: dimensions[1],
                height: dimensions[2],
            }
        })
        .collect();

    let total_area_required: i32 = presents.iter().map(|p| p.surface_area_and_extra()).sum();
    let total_ribbon_required: i32 = presents.iter().map(|p| p.ribbon_needed()).sum();

    println!("Total area to purchase: {}", total_area_required);
    println!("Total ribbon to purchase: {}", total_ribbon_required);
}
