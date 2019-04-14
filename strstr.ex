File.stream!("tmp/logs.log")
|> Stream.map(&(String.trim(&1) |> String.split("OK ", parts: 2) |> Enum.at(1)))
|> Enum.filter(&(&1 != nil))
|> IO.inspect()
