perfs = []

File.foreach("tmp/logs.log") do |line|
  if line.length > 4
    perf = line.split("OK ", 2)
    perfs << perf[1].strip! if perf.length == 2
  end
end

p perfs[0..40]
