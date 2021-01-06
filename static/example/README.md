# Chrome Remote applet

这是 JumpServer Remote applet 的一个 example

## 关于 i18n 使用

```
import jinja2, yaml
from jinja2 import Environment, FileSystemLoader

f = open('meta.yml')
i18n_data = yaml.safe_load(f)['i18n']
env = Environment(loader=FileSystemLoader('.'))

def i18n(value):
return i18n_data.get('cn', {}).get(value, value)
env.filters['i18n'] = i18n


template = env.get_template('meta.yml')
d2 = template.render()
print(d2)
yaml.loads(d2)

import json
print(template.render())
d2 = template.render()
yaml.safe_load(d2)
```