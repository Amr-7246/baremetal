import unittest 
from problems import split_str

class spliterTest(unittest.TestCase):
    def test_1(self):
        self.assertEqual( split_str("amr"), "Here is You splited string ['am', 'r_']" )
    def test_2(self):
        self.assertEqual( split_str("a"), "Here is You splited string ['a_']" )
        
if __name__ == "__main__" :
    unittest.main()
