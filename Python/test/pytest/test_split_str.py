from Python.problems import split_str

def test_1():
    assert split_str("amr") == 'Here is You splited string ["am", "r_"]' 
def test_2():
    assert split_str("a") == 'Here is You splited string ["a_"]'
