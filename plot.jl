using Plots 
using Dates
using Statistics
using Measures

function load(path::String)
    # open file, read data line-by-line
    x = Vector{Float64}()
    y = Vector{Float64}()
    open(path) do file
        for (i,ln) in enumerate(eachline(file))
            if i%2 == 0 
                append!(y, parse(Float64, ln))
            else
                append!(x, parse(Float64, ln))
            end
        end
    end
x,y
end

function plt(x::Vector{Float64},y::Vector{Float64})
    # Example Unix time series and corresponding data
    unix_time_series = y  # Your Unix time series data in float64 format
    data_values = x       # Your corresponding data values

    data_values = data_values/mean(data_values)
    # Convert Unix time to DateTime
    datetime_series = map(unix2datetime, unix_time_series)

    # Calculate a reasonable tick interval based on the data range
    t = map(unix2datetime, LinRange(unix_time_series[1],unix_time_series[end],15))
    time_ticks = Dates.format.(t,"dd,mm,yy")

    # Plot the data with custom tick rate
    plot_object = plot(datetime_series, data_values, xticks = (t, time_ticks), xrotation = 45, fmt = :auto, linewidth=2, title="Relative local B-field", label="B-field", dpi=300, margin=3mm)
    xlabel!("Time")
    ylabel!("relative field strength")

    # Display and save the plot

    savefig(plot_object, "plot.png")  # Adjust the filename and format as needed
    return 0
end

x,y = load(ARGS[1])
plt(x,y)
