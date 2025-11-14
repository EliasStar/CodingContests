use std::{error::Error, fs};

static INPUT_PATHS: &[&str] = &[
    "level2/level2_0_example",
    "level2/level2_1_small",
    "level2/level2_2_large",
];

struct Input {
    sequence_count: u32,
    sequences: Vec<Vec<i32>>,
}

struct Output {
    pos_and_times: Vec<(i32, u32)>,
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

    let sequence_count = lines.next().ok_or("malformed")?.parse::<u32>()?;
    let mut sequences = Vec::with_capacity(sequence_count as usize);

    for line in lines {
        let r: Result<Vec<i32>, _> = line.split_ascii_whitespace().map(str::parse).collect();
        sequences.push(r?);
    }

    Ok(Input {
        sequence_count,
        sequences,
    })
}

fn compute_pos_and_time(input: &Input) -> Output {
    let mut pos_and_times = Vec::with_capacity(input.sequence_count as usize);

    for seq in &input.sequences {
        let mut pos = 0i32;
        let mut time = 0u32;

        for step in seq {
            pos += if *step > 0 {
                1
            } else if *step < 0 {
                -1
            } else {
                0
            };

            time += if *step == 0 { 1 } else { step.unsigned_abs() };
        }

        pos_and_times.push((pos, time));
    }

    Output { pos_and_times }
}

fn write_output(path: &str, output: &Output) -> Result<(), Box<dyn Error>> {
    fs::write(
        path.to_owned() + ".out",
        output
            .pos_and_times
            .iter()
            .map(|(pos, time)| format!("{pos} {time}"))
            .collect::<Vec<_>>()
            .join("\n"),
    )?;
    Ok(())
}
