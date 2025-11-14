use std::{error::Error, fs};

static INPUT_PATHS: &[&str] = &[
    "level3/level3_0_example",
    "level3/level3_1_small",
    "level3/level3_2_large",
];

const MIN_PACE: u32 = 5;
const FULL_ACCEL_LEN: u32 = 2 * MIN_PACE - 1;

struct Input {
    case_count: u32,
    cases: Vec<(i32, u32)>,
}

struct Output {
    sequences: Vec<Vec<i32>>,
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
        let mut split = line.split_ascii_whitespace();

        let target_pos = split.next().ok_or("case malformed")?.parse::<i32>()?;
        let time_limit = split.next().ok_or("case malformed")?.parse::<u32>()?;

        cases.push((target_pos, time_limit));
    }

    Ok(Input { case_count, cases })
}

fn compute_pos_and_time(input: &Input) -> Output {
    let mut sequences = Vec::with_capacity(input.case_count as usize);

    for (target_pos, _time_limit) in &input.cases {
        let mut sequence = vec![0];

        sequence.extend(
            generate_accel_decel_vec(target_pos.unsigned_abs())
                .into_iter()
                .map(|x| {
                    if *target_pos < 0 {
                        -(x as i32)
                    } else {
                        x as i32
                    }
                }),
        );

        sequence.push(0);
        sequences.push(sequence);
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
            .map(|seq| seq.iter().map(i32::to_string).collect::<Vec<_>>().join(" "))
            .collect::<Vec<_>>()
            .join("\n"),
    )?;
    Ok(())
}
