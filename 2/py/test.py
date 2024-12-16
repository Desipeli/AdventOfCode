import unittest
from functions import is_safe

class TestIsSafe(unittest.TestCase):
    def test_correct_asc(self):
        input = ["1", "2", "4", "6", "9"]
        self.assertEqual(is_safe(input), True)
    
    def test_correct_desc(self):
        input = ["4","3","2","1"]
        self.assertEqual(is_safe(input), True)

    def test_same_subsequent_numbers_must_fail(self):
        input = ["1", "2", "2", "3"]
        self.assertEqual(is_safe(input), False)
    
    def test_asc_over_3_fails(self):
        input = ["1", "2", "6"]
        self.assertEqual(is_safe(input), False)
    
    def test_desc_over_3_fails(self):
        input = ["6", "2", "1"]
        self.assertEqual(is_safe(input), False)
    
    def test_asc_after_desc_fails(self):
        input = ["3", "2", "1", "2"]
        self.assertEqual(is_safe(input), False)

    def test_desc_after_asc_fails(self):
        input = ["1", "2", "3", "1"]
        self.assertEqual(is_safe(input), False)
        
if __name__ == '__main__':
    unittest.main()