# Ramp Challenge

Thanks for applying to Ramp. Solve this CTF[1] challenge and add the result to your application.

We recommend opening this file with a Markdown viewer. (https://www.google.com/search?q=markdown+viewer)

## Instructions

1. Open this [link](https://tns4lpgmziiypnxxzel5ss5nyu0nftol.lambda-url.us-east-1.on.aws/challenge)
2. Find a hidden URL within the HTML
   - Each character of the URL is given by this DOM tree, in this specific order. You need to find (in order) all of the occurrences and join them to get the link.
   - The asterisk **(\*)** is a wildcard representing zero or more characters that can be present in the string. These characters are irrelevant to the result and should be ignored.
   - There can be zero or more DOM nodes between each valid tag. These nodes are irrelevant to the result.
   - Any additional attribute that doesn't interfere with the described pattern can be safely ignored.

Pattern of the DOM tree for each valid character of the URL

```html
<code data-class="23*">
  <div data-tag="*93">
    <span data-id="*21*">
      <i class="char" value="VALID_CHARACTER"></i>
    </span>
  </div>
</code>
```

(_To validate this step, you should be able to open the URL and get an English word. This means you have captured the flag!_ 🥳)
