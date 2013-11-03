cfn-summarize
======

**Summarizes a CloudFormation template.**

This prints a list of the things in a CloudFormation template:

    cfn-summarize [-a] [-s] template
    
The ```-a``` flag will show acme addresses to the item. Probably only useful to [@drocamor](https://github.com/drocamor).

The ```-s``` flag will show the command to run [cfn-show](https://github.com/controlgroup/cfn-show) to detail the item.