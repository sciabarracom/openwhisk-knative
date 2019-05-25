FROM julia:1.1.0
ADD setup.jl .
RUN julia setup.jl
CMD julia -e 'using OpenWhiskKnative; start()'
