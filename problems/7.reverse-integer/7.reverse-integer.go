
package main

import (
	"fmt"
	"strconv"
)

/**
 * <p>Given a signed 32-bit integer <code>x</code>, return <code>x</code><em> with its digits reversed</em>. If reversing <code>x</code> causes the value to go outside the signed 32-bit integer range <code>[-2<sup>31</sup>, 2<sup>31</sup> - 1]</code>, then return <code>0</code>.</p>

<p><strong>Assume the environment does not allow you to store 64-bit integers (signed or unsigned).</strong></p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> x = 123
<strong>Output:</strong> 321
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> x = -123
<strong>Output:</strong> -321
</pre><p><strong>Example 3:</strong></p>
<pre><strong>Input:</strong> x = 120
<strong>Output:</strong> 21
</pre><p><strong>Example 4:</strong></p>
<pre><strong>Input:</strong> x = 0
<strong>Output:</strong> 0
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-2<sup>31</sup> &lt;= x &lt;= 2<sup>31</sup> - 1</code></li>
</ul>

**/
/**
 * 123
**/
func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x *= -1
	}

	runes := []rune(fmt.Sprintf("%d", x))
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j -1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	ret, err := strconv.Atoi(string(runes))
	if err != nil {
		return 0
	}

	if negative {
		ret = ret * -1

	}

	if ret < (2 << 30) * -1 || (2 << 30)  - 1 < ret {
		return 0
	}
	return ret
}
