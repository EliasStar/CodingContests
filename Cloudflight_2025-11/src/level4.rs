use std::{error::Error, fs};

static INPUT_PATHS: &[&str] = &[
    "level4/level4_0_example",
    "level4/level4_1_small",
    "level4/level4_2_large",
];

const MIN_PACE: u32 = 5;
const FULL_ACCEL_LEN: u32 = 2 * MIN_PACE - 1;

struct Input {
    case_count: u32,
    cases: Vec<(i32, i32)>,
}

struct Output {
    sequences: Vec<(Vec<i32>, Vec<i32>)>,
}

pub fn run() -> Result<(), Box<dyn Error>> {
    for path in INPUT_PATHS {
        let input = read_input(path)?;
        let output = compute_pos_and_time(&input);
        write_output(path, &output)?;
    }
    Ok(())
}

fn read_input(path: &str) -> Result<Input, Box<dyn Error>> {
    let input_str = fs::read_to_string(path.to_owned() + ".in")?;
    let mut lines = input_str.lines();

    let case_count = lines.next().ok_or("malformed")?.parse::<u32>()?;
    let mut cases = Vec::with_capacity(case_count as usize);

    for line in lines {
        let (pos_x, pos_y) = line
            .split_ascii_whitespace()
            .next()
            .ok_or("case malformed")?
            .split_once(",")
            .ok_or("case malformed")?;

        cases.push((pos_x.parse::<i32>()?, pos_y.parse::<i32>()?));
    }

    Ok(Input { case_count, cases })
}

fn compute_pos_and_time(input: &Input) -> Output {
    let mut sequences = Vec::with_capacity(input.case_count as usize);

    for (pos_x, pos_y) in &input.cases {
        let mut sequence_x = vec![0];
        let mut sequence_y = vec![0];

        sequence_x.extend(
            generate_accel_decel_vec(pos_x.unsigned_abs())
                .into_iter()
                .map(|x| if *pos_x < 0 { -(x as i32) } else { x as i32 }),
        );

        sequence_y.extend(
            generate_accel_decel_vec(pos_y.unsigned_abs())
                .into_iter()
                .map(|x| if *pos_y < 0 { -(x as i32) } else { x as i32 }),
        );

        sequence_x.push(0);
        sequence_y.push(0);
        sequences.push((sequence_x, sequence_y));
    }

    Output { sequences }
}

fn generate_accel_decel_vec(length: u32) -> Vec<u32> {
    if length == 0 {
        return vec![];
    }

    if length >= FULL_ACCEL_LEN {
        let mut vec: Vec<u32> = (1..=MIN_PACE).rev().collect();
        vec.extend(vec![1; (length - FULL_ACCEL_LEN) as usize]);
        vec.extend(2..=MIN_PACE);
        vec
    } else {
        let len_accel = (length as f64 / 2.0).ceil() as u32;
        let len_decel = length - len_accel;

        let mut vec: Vec<u32> = (MIN_PACE - len_accel + 1..=MIN_PACE).rev().collect();
        vec.extend((MIN_PACE - len_decel + 1..=MIN_PACE).collect::<Vec<u32>>());
        vec
    }
}

fn write_output(path: &str, output: &Output) -> Result<(), Box<dyn Error>> {
    fs::write(
        path.to_owned() + ".out",
        output
            .sequences
            .iter()
            .map(|(seq_x, seq_y)| {
                let seq_x_str = seq_x
                    .iter()
                    .map(i32::to_string)
                    .collect::<Vec<_>>()
                    .join(" ");
                let seq_y_str = seq_y
                    .iter()
                    .map(i32::to_string)
                    .collect::<Vec<_>>()
                    .join(" ");
                format!("{seq_x_str}\n{seq_y_str}\n")
            })
            .collect::<Vec<_>>()
            .join("\n"),
    )?;
    Ok(())
}
