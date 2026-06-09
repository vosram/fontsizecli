# Fontsize CLI

This CLI tool helps to quickly create a set a font sizes for web design or custom graphic designs using custom ratios. For web design, this tool can create desktop and mobile versions using different base sizes (the size of the standard p element), and ratios. All pixel values are rounded to the nearest pixel, i.e. `35.6` -> `36px`.

## Installation

You can install this tool with the following command:

```bash
go install github.com/vosram/fontsizecli
```

## Usage

There are 3 subcommands: `list`, `web` and `custom`.

### List

This command prints out common ratios you can use to create these set of font sizes. This is a handy command if you're unsure of what ratio you should use.

Example:

```bash
fontsizecli list
```

### Web

This command creates font size sets for desktop and mobile designs. by default it creates the a size set for both desktop and mobile using the same base size and ratio. with added flags, you can create separate size sets for mobile design.

**Basic usage:**

This command would create a fontsize set for a base font of `16px` and a ratio of `1.25`.

```bash
# fontsizecli web <base-size-integer> <ratio-float64>
fontsizecli web 16 1.25
```

**Output:**

```text
Desktop/Mobile Sizes:
-  h1: 61px
-  h2: 49px
-  h3: 39px
-  h4: 31px
-  h5: 25px
-  h6: 20px
-   p: 16px
- smp: 13px
- xsp: 10px
```

**Desktop and Mobile Sizes Example:**

```bash
# fontsizecli web <desktop-base-size> <desktop-ratio> --mob-base <mobile-base-size> --mob-ratio <mobile-ratio>

fontsize web 16 1.25 --mob-base 18 --mob-ratio 1.2
```

**Output:**

```text
Desktop Sizes:
-  h1: 61px
-  h2: 49px
-  h3: 39px
-  h4: 31px
-  h5: 25px
-  h6: 20px
-   p: 16px
- smp: 13px
- xsp: 10px

Mobile Sizes:
-  h1: 53px
-  h2: 44px
-  h3: 37px
-  h4: 31px
-  h5: 26px
-  h6: 22px
-   p: 18px
- smp: 15px
- xsp: 13px
```

#### Web flags

`--mob-base <base-size>` will generate the same sizes for h1-h6, p, smp, xsp for mobile starting at this base size.

`--mob-ratio` will generate the mobile sizes with this ratio. Default is `1.2`.

`-o <filename>` exports a txt file to current working directory. The extension `.txt` is added automatically.

### Custom

This command will create a set of font sizes in cases that go outside of the normal web design workflow. You can give it a base font size either as the highest value or lowest value, and a ratio and how many steps you want.

#### Usage

**Example:**

```bash
# fontsize custom <h || l> <base-font-size> <ratio> <steps>

fontsizecli custom h 140 1.25 5
```

**Output:**

```text
1 - 140px
2 - 112px
3 - 90px
4 - 72px
5 - 58px
```

**Example 2:**

```bash
# fontsize custom <h || l> <base-font-size> <ratio> <steps>

fontsizecli custom l 20 1.25 5
```

**Output:**

```text
1 - 20px
2 - 25px
3 - 31px
4 - 39px
5 - 49px
```

#### Custom Flags

This command only has one flag `-o <filename>` that writes a file with generated results to a file in the working directory.
