use std::fs::File;
use std::io::{BufReader, BufRead};
use std::io::Error;

fn main() {
    match run() {
      Ok(_) => println!("done!"),
      _error => println!("uhoh!")
    }
}

fn run() -> Result<(), Error> {
  let mut lines: Vec<String> = vec![];

  let file = File::open("tmp/logs.log")?;
  let buffered = BufReader::new(file);

  for line in buffered.lines() {
      let ln = line?;
      if ln.chars().count() > 3 {
          if ln.contains("OK db=") {
            let result: Vec<&str> = ln.split("OK ").collect();

            let st = String::from(result[1]);

            lines.push(st);    
          }
      }
  }

  println!("{}\n{}\n{}", lines[38], lines[39], lines[40]);

  Ok(())
}
