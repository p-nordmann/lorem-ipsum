# lorem-ipsum
A simple _lorem ipsum_ generator based on [prefix-forest](https://github.com/p-nordmann/prefix-forest).

## Introduction
This repository provides a simple command to generate _lorem ipsum_ like paragraphs.

It uses the _Liber Primus_, copied from [thelatinlibrary.com](https://www.thelatinlibrary.com/cicero/fin.shtml) and tries to generate similar text based on `<5`-gram statistics.

## Installation and usage

Right now, the CLI has to be run from sources:
```
git clone https://github.com/p-nordmann/lorem-ipsum
cd lorem-ipsum
go run ./cmd/glorips --paragraphs 2
```
The commands above can result in the following, e.g.:
>Sunt anima est? quod ipso patiam nisi inito insita dissimos satis magna affecta, quae plus non possum, quae praeci te, quod omnium rerum erit. voluit, magnum dici poena legerem, ut postra metus gratio. nam et pertatibus in ea faciam: declinam nostrum venire poster noster existincto. locatus decordamur mortis, idem, quidem motu venanditur Democritudinem discipitur, alterae sit iustitudinem est enim quem atque liberiora inime praetero quam cum suo, genus vitae motus est cuius disputavisse. -- Filius de Graecum omnes mi Albucius, tota referret reicienda non postra ipsius susceperii ne efficeret? Alii partes expectas ex amicorum, cum ardore ea, quae quae perpetulet non modius. Alienum in bona.
>
> Sed existimant usque potest, hominus a studines autem, ubi corpora interrogari. et partitudo sine propter arridens secunda, fruuntur. illa, qui putamur est propter placete veterear, numquam fugiendide dicere maestatus nimis voluntatem physicine sapientibus nullam et per uterque studiis satiam, neque inflammati multas, quae fugiendis vitatum tantam tanta ab inliberat ille meliae pareat eligi et corribimus, Brutus et apti est a provincidunt, summam voluptatem, inquam, tamque se iudicium requirenda nec impetu omnium animi et quam quidem ipsum, cum in dolor autem cursu ad voluptas non potest, sed ut vult, ea ratio tractantum, quod sit. re decordationis reprehensionibus saepe everte ne Graecisse?

## Roadmap
- Control length of paragraphs by boosting probability of `\n` token as paragraphs get longer.
- Clean _Liber Primus_ from special characters.
- Provide the option to start generated text with _Lorem ipsum dolor..._.

## License
MIT License. See [LICENSE](LICENSE) file for more details.
