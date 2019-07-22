function retry(f::Function, max_iter, sleep_for=1)
    for i in 1:max_iter
        if f()
            return true
        end
        @info "not ready yet, retring in $sleep_for second(s)"
        sleep(sleep_for)
    end
    @warn "max retries ($max_iter) reached, exiting"
    return false 
end

function run_log(cmd::Cmd) 
    @info join(cmd.exec, " ")
    run(cmd)
end

function read_log(cmd::Cmd, args...)
    @info join(cmd.exec, " ")
    return read(cmd, String)
end
