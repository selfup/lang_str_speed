perfs = []

File.foreach("tmp/logs.log") do |line|
  line.strip!

  if line.length > 0
    perf = line.split("OK ", 2)
    perfs << perf[1] if perf.length == 2
  end
end
