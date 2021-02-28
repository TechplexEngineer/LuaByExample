<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Lua by Example: For</title>
    <link rel=stylesheet href="site.css">
  </head>
  <script>
      onkeydown = (e) => {
          
          if (e.key == "ArrowLeft") {
              window.location.href = 'variables';
          }
          
          
          if (e.key == "ArrowRight") {
              window.location.href = 'while';
          }
          
      }
  </script>
  <body>
    <div class="example" id="for">
      <h2><a href="./">Lua by Example</a>: For</h2>
      
      <table>
        
        <tr>
          <td class="docs">
            <p><code>for</code> is one of Lua&rsquo;s looping constructs. Here are
some basic types of <code>for</code> loops.</p>

          </td>
          <td class="code empty leading">
            
          
          </td>
        </tr>
        
        <tr>
          <td class="docs">
            <p>A three expression loop is called a <em>numeric for</em> loop
i=1 is the initial condition, 10 is the end condition, 1 is the step increment
the third expression is optional, when omitted 1 is used as the step.</p>

          </td>
          <td class="code leading">
            
          <pre class="chroma">
<span class="kr">for</span> <span class="n">i</span><span class="o">=</span><span class="mi">1</span><span class="p">,</span> <span class="mi">10</span><span class="p">,</span> <span class="mi">1</span> <span class="kr">do</span>
    <span class="n">print</span><span class="p">(</span><span class="n">i</span><span class="p">)</span>
<span class="kr">end</span></pre>
          </td>
        </tr>
        
        <tr>
          <td class="docs">
            <p>Note the variable declared in the loop is local to the body of the loop
despite the lack of the local keyword.</p>

          </td>
          <td class="code leading">
            
          <pre class="chroma">
<span class="n">print</span><span class="p">(</span><span class="n">i</span><span class="p">)</span> <span class="c1">-- will print nil</span></pre>
          </td>
        </tr>
        
        <tr>
          <td class="docs">
            <p><code>break</code> can be used to end the loop early</p>

          </td>
          <td class="code leading">
            
          <pre class="chroma">
<span class="kr">for</span> <span class="n">i</span><span class="o">=</span><span class="mi">1</span><span class="p">,</span> <span class="mi">10</span><span class="p">,</span> <span class="mi">1</span> <span class="kr">do</span>
    <span class="n">print</span><span class="p">(</span><span class="n">i</span><span class="p">)</span>
    <span class="kr">if</span> <span class="n">i</span> <span class="o">==</span> <span class="mi">5</span> <span class="kr">then</span>
        <span class="kr">break</span>
    <span class="kr">end</span>
<span class="kr">end</span></pre>
          </td>
        </tr>
        
        <tr>
          <td class="docs">
            <p>A <em>generic for</em> loop supports traversing elements returned from an iterator
print all values of array <code>arr</code></p>

          </td>
          <td class="code leading">
            
          <pre class="chroma">
<span class="kr">for</span> <span class="n">i</span><span class="p">,</span><span class="n">v</span> <span class="kr">in</span> <span class="n">ipairs</span><span class="p">(</span><span class="n">arr</span><span class="p">)</span> <span class="kr">do</span>
    <span class="n">print</span><span class="p">(</span><span class="n">v</span><span class="p">)</span>
<span class="kr">end</span></pre>
          </td>
        </tr>
        
        <tr>
          <td class="docs">
            <p>Lua does not have a <code>continue</code> construct. A workaround is to use goto.
prints odd numbers in [|1,10|]</p>

          </td>
          <td class="code">
            
          <pre class="chroma">
<span class="kr">for</span> <span class="n">i</span><span class="o">=</span><span class="mi">1</span><span class="p">,</span><span class="mi">10</span> <span class="kr">do</span>
  <span class="kr">if</span> <span class="n">i</span> <span class="o">%</span> <span class="mi">2</span> <span class="o">==</span> <span class="mi">0</span> <span class="kr">then</span> <span class="kr">goto</span> <span class="nl">continue</span> <span class="kr">end</span>
  <span class="n">print</span><span class="p">(</span><span class="n">i</span><span class="p">)</span>
  <span class="p">::</span><span class="nl">continue</span><span class="p">::</span>
<span class="kr">end</span></pre>
          </td>
        </tr>
        
      </table>
      
      <table>
        
        <tr>
          <td class="docs">
            
          </td>
          <td class="code leading">
            
          <pre class="chroma"><span class="gp">$</span> lua for.lua</pre>
          </td>
        </tr>
        
        <tr>
          <td class="docs">
            <p>We&rsquo;ll see some other <code>for</code> forms later when we look at
<code>range</code> statements, channels, and other data
structures.</p>

          </td>
          <td class="code empty">
            
          
          </td>
        </tr>
        
      </table>
      
      
      <p class="next">
        Next example: <a href="while">While</a>.
      </p>
      
      <p class="footer">
        by <a href="https://techplexlabs.com">Blake Bourque</a> |
        <a href="https://github.com/TechplexEngineer/LuaByExample/blob/master/examples/for">source</a> |
        <a href="https://github.com/TechplexEngineer/LuaByExample#license">license</a>
      </p>
    </div>
    <script>
      var codeLines = [];
      codeLines.push('');codeLines.push('');codeLines.push('');codeLines.push('');codeLines.push('');codeLines.push('');codeLines.push('');codeLines.push('');
    </script>
    <script src="site.js" async></script>
  </body>
</html>
