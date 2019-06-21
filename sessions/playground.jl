cmd = """wc -l <EOF
    a
    b
    c
    EOF
    """

l = IOBuffer("""hello
    world
    """)

 m = replace("""first
       second
       third
       """, '\n' => ' ')
