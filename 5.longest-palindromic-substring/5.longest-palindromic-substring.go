
package main

/**
 * <p>Given a string <code>s</code>, return&nbsp;<em>the longest palindromic substring</em> in <code>s</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;babad&quot;
<strong>Output:</strong> &quot;bab&quot;
<strong>Note:</strong> &quot;aba&quot; is also a valid answer.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;cbbd&quot;
<strong>Output:</strong> &quot;bb&quot;
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;a&quot;
<strong>Output:</strong> &quot;a&quot;
</pre>

<p><strong>Example 4:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;ac&quot;
<strong>Output:</strong> &quot;a&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> consist of only digits and English letters.</li>
</ul>

**/
/**
 * "babad"
**/
func longestPalindrome(s string) string {
	ret := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			substrLength := j - i + 1
			isPalindrome := true
			for k := 0; k < substrLength / 2; k++ {
				if s[i + k] != s[j - k] {
					isPalindrome = false
					break
				}
			}

			if isPalindrome {
				if len(ret) < j - i + 1{
					ret = s[i:(j + 1)]
				}
				break
			}
		}
	}
	return ret
}
