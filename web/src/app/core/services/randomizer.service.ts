import { Injectable } from '@angular/core'; 

import { AnswerBoxModel } from '../models/answer-box.model';
import { QuestionBoxEntryModel } from '../models/question-box-entry.model';
import { AnswerBoxEntryModel } from '../models/answer-box-entry.model';

@Injectable()
export class RandomizerService {
  randomPositiveWords: [string];

  constructor() { }

  getRandomWordCloudEntriesWithDefault(defaultValue: string, count:number):[[string,number]]{
    let entries: [[string,number]] = [[defaultValue,110]];

    for(let i = 0; i<count; i++){
      let randomSize = this.getRandomMultipleOf10();
      let randomPositiveWord = this.getRandomBaseballTeams();
      entries.push([randomPositiveWord,randomSize]);
    } 
    
    return entries;
  }

  getRandomQuestionBoxEntries(count:number):QuestionBoxEntryModel[]{
    let entries: QuestionBoxEntryModel[] = [];

    let questions = [
      "What are our goals for next period?",
      "What can we improve on in the future?",
      "How can we provide more value to our clients?",
      "How can we be more transparent with internal changes?",
      "How can we adapt to reorgs more easily?",
      "How is the company doing in relation to our competitors?",
      "How can we engage with our customers more easily?",
      "When will we be in the top 100 companies to work for?",
      "What did we do for our community this past period?",
      "Why are we not prioritizing customer feedback?"
    ]

    count = questions.length > count? count: questions.length;

    for(let i = 0; i<=count; i++){
      let entry:QuestionBoxEntryModel = { Question: questions[i], Downvotes:0, Upvotes:1 }
      entries.push(entry);
    }

    return entries;
  }

  getRandomAnswerBoxEntries(count:number):AnswerBoxEntryModel[]{
    let max = this.getRandomNumberByMaxValue(count);
    let entries:AnswerBoxEntryModel[] = []; 
    for(let i = 0; i<=max; i++){
      let entry:AnswerBoxEntryModel = {
        Entry: this.getRandomCupcakeSentence(),
        IsHidden: false,
      };
      entries.push(entry);
    }

    return entries;
  }

  getRandomVotesBoxEntries(count:number):number{
    let entriesCount = this.getRandomNumberByMaxValue(count); 
    return entriesCount;
  }

  private getRandomCupcakeSentence():string{
    let cupcakeSentence:string;
    let answers = ["Candy canes sugar plum chupa chups sesame snaps powder marzipan candy sesame snaps.","Icing gingerbread cheesecake brownie jelly beans jelly sweet roll.","Chupa chups sugar plum marshmallow jelly beans cheesecake.","Chocolate cake croissant jelly beans.","Apple pie lollipop cake danish bonbon.","Chocolate bar gummies halvah sugar plum biscuit halvah bear claw dessert.","Jelly-o jelly liquorice jelly-o powder ice cream lemon drops.","Soufflé tiramisu macaroon dragée jujubes tart.","Danish macaroon powder pastry liquorice.","Macaroon pudding marzipan donut liquorice pudding jujubes lemon drops dragée.","Pudding lollipop chocolate ice cream marshmallow danish pastry.","Chocolate bar cupcake jelly beans biscuit oat cake icing.","Wafer apple pie cookie macaroon croissant tootsie roll cake.","Liquorice jujubes brownie.","Marzipan gummies cake apple pie caramels oat cake pastry.","Tootsie roll carrot cake oat cake bonbon jujubes chupa chups.","Oat cake gummies cheesecake.","Biscuit gummi bears chocolate cake chocolate bar bear claw soufflé sugar plum donut powder.","Jelly-o cotton candy gummies caramels sesame snaps caramels carrot cake.","Tart gummies croissant liquorice chocolate muffin chocolate bar.","Pudding jelly sugar plum cake dessert halvah liquorice sweet danish.","Croissant gummies pudding dessert chocolate bar chocolate chupa chups dessert.","Tootsie roll cake caramels cake.","Macaroon cookie candy canes fruitcake halvah candy.","Brownie lemon drops carrot cake marzipan ice cream cupcake lemon drops.","Jelly chocolate bar tiramisu pastry topping croissant chupa chups danish.","Halvah topping cake cake cupcake marzipan tiramisu.","Lollipop halvah croissant icing lemon drops halvah.","Jelly beans jujubes gingerbread wafer sweet croissant.","Muffin jelly-o cake apple pie sesame snaps cookie jelly.","Caramels chupa chups dragée marshmallow icing.","Pie bonbon jelly-o chupa chups.","Fruitcake liquorice apple pie jelly-o candy cookie sweet roll sweet chupa chups.","Chupa chups bonbon fruitcake danish marshmallow chocolate.","Dessert tart pie gingerbread.","Dragée jelly beans donut gingerbread icing cheesecake.","Lollipop marshmallow bear claw gingerbread cookie marshmallow.","Jelly-o cake topping bear claw chocolate bar.","Soufflé caramels gummies.","Pie fruitcake tart gingerbread jelly-o lemon drops pudding ice cream tiramisu.","Danish biscuit chocolate bar candy.","Lollipop topping croissant jelly beans candy cake bear claw.","Sweet roll lemon drops bonbon carrot cake.","Candy canes brownie candy gummies danish marshmallow cake tiramisu.","Cookie muffin candy icing sweet sesame snaps liquorice brownie dessert.","Cotton candy wafer dessert tart.","Gingerbread brownie fruitcake.","Powder pudding oat cake donut cotton candy cotton candy cupcake icing donut.","Carrot cake powder powder danish muffin powder lollipop soufflé muffin.","Cake chocolate bar powder sesame snaps biscuit cake.","Jelly-o cupcake powder.","Pudding cookie fruitcake powder tiramisu.","Jujubes soufflé apple pie toffee cheesecake muffin dessert.","Bear claw jujubes cookie.","Carrot cake wafer ice cream caramels jujubes liquorice candy canes soufflé.","Cookie cotton candy cake.","Candy canes marzipan marzipan sweet roll cake chocolate bar fruitcake toffee cheesecake.","Cheesecake muffin sweet powder jelly-o sesame snaps marzipan.","Marzipan sugar plum powder jelly marzipan brownie brownie jujubes.","Brownie tootsie roll sweet roll tootsie roll jelly beans ice cream sugar plum candy sugar plum.","Chupa chups muffin pie jelly pastry dragée sesame snaps jelly-o topping.","Croissant bonbon cookie cupcake lemon drops gummies jelly beans sweet roll.","Oat cake bear claw soufflé gummies oat cake.","Gingerbread bonbon bear claw pastry topping chocolate bar halvah gingerbread.","Chocolate cake cookie oat cake pie apple pie sesame snaps chocolate bar.","Marzipan jelly beans brownie dessert tiramisu.","Toffee powder cupcake topping candy chocolate biscuit.","Tart caramels jujubes icing sweet roll.","Bonbon danish soufflé biscuit jelly beans.","Bear claw tootsie roll cake liquorice.","Ice cream cake oat cake cookie topping croissant cake chocolate bar.","Liquorice sesame snaps wafer.","Chupa chups jelly-o cheesecake gummi bears apple pie sugar plum biscuit.","Gummi bears danish ice cream tiramisu tiramisu muffin.","Biscuit macaroon sweet roll.","Donut icing halvah gingerbread marzipan.","Chocolate bar caramels sweet tart chocolate cake marzipan cake chupa chups pie.","Brownie biscuit icing lemon drops sweet roll pudding tiramisu.","Caramels bonbon jelly-o jujubes cake cake lemon drops.","Sweet wafer cupcake dragée.","Cake pastry biscuit.","Sweet roll candy canes macaroon.","Tootsie roll carrot cake gummi bears macaroon jujubes jelly donut candy.","Cake brownie ice cream.","Gummi bears donut tiramisu toffee.","Sesame snaps cupcake apple pie oat cake oat cake caramels jelly-o.","Croissant gingerbread chocolate caramels cake lollipop halvah cotton candy marshmallow.","Icing icing lemon drops bear claw dragée.","Chocolate cookie jelly donut wafer sweet roll.","Cake halvah soufflé pie jelly icing cake pie fruitcake.","Dessert gingerbread bear claw oat cake cotton candy pie cake.","Cookie danish caramels gummi bears jelly-o lemon drops macaroon danish.","Apple pie topping chocolate bar candy canes biscuit jelly powder toffee croissant.","Dragée jelly brownie sweet macaroon sweet sweet roll.","Soufflé cookie muffin muffin gingerbread.","Tootsie roll marzipan gingerbread candy canes powder tiramisu chupa chups cake.","Carrot cake cheesecake candy apple pie jujubes chupa chups.","Cake pie gummi bears biscuit muffin chupa chups pudding sweet roll apple pie.","Sweet sweet tiramisu jelly tootsie roll.","Carrot cake icing lemon drops cookie chocolate.","Dessert donut tiramisu.","Muffin pastry croissant.","Jelly-o cake liquorice muffin bear claw.","Tart cake cotton candy muffin bear claw muffin.","Donut liquorice halvah biscuit gummies.","Bonbon oat cake sweet roll icing lemon drops.","Soufflé apple pie sweet roll.","Liquorice sweet roll jelly tiramisu bonbon donut danish sweet sugar plum.","Liquorice candy dragée candy canes.","Chocolate cake halvah candy canes cupcake tootsie roll pastry.","Gingerbread macaroon cotton candy chocolate bar.","Lemon drops icing marshmallow gummi bears brownie cake gummi bears jelly beans.","Caramels marzipan jelly-o sweet roll danish fruitcake gummies cheesecake.","Sesame snaps gingerbread macaroon soufflé.","Chupa chups jujubes gummies.","Topping icing sesame snaps marshmallow toffee wafer liquorice chocolate bar.","Oat cake candy candy fruitcake biscuit.","Gummies tiramisu chupa chups icing carrot cake cotton candy gummi bears cheesecake.","Ice cream marshmallow cheesecake pudding jelly-o biscuit cake toffee gingerbread.","Chocolate bar muffin toffee cookie donut.","Cheesecake cupcake sesame snaps powder tart toffee muffin.","Brownie halvah chocolate topping soufflé jelly jujubes topping.","Toffee carrot cake marshmallow sweet roll pie bonbon.","Toffee jelly-o liquorice gummies bear claw fruitcake dragée.","Chocolate bar gummi bears pie sesame snaps gummi bears jujubes.","Ice cream oat cake bear claw lemon drops.","Jelly beans fruitcake candy chocolate cake.","Cookie ice cream toffee tiramisu cake dragée.","Lollipop biscuit lemon drops marshmallow tootsie roll fruitcake oat cake chocolate cake.","Muffin fruitcake cookie gummi bears tootsie roll chocolate.","Carrot cake cookie pudding marzipan tiramisu.","Pastry sweet candy canes tootsie roll fruitcake chocolate cake soufflé."];
    let randomIndex = this.getRandomIndexForArray(answers);
    cupcakeSentence = answers[randomIndex];
    return cupcakeSentence;
  }

  private getRandomPositiveWord():string{
    let positiveWord:string;
    let positiveWords = ["Absolutely","Abundant","Accept","Acclaimed","Accomplishment","Achievement","Action","Active","Activist","Acumen","Adjust","Admire","Adopt","Adorable","Adored","Adventure","Affirmation","Affirmative","Affluent","Agree","Airy","Alive","Alliance","Ally","Alter","Amaze","Amity","Animated","Answer","Appreciation","Approve","Aptitude","Artistic","Assertive","Astonish","Astounding","Astute","Attractive","Authentic","Basic","Beaming","Beautiful","Believe","Benefactor","Benefit","Bighearted","Blessed","Bliss","Bloom","Bountiful","Bounty","Brave","Bright","Brilliant","Bubbly","Bunch","Burgeon","Calm","Care","Celebrate","Certain","Change","Character","Charitable","Charming","Cheer","Cherish","Clarity","Classy","Clean","Clever","Closeness","Commend","Companionship","Complete","Comradeship","Confident","Connect","Connected","Constant","Content","Conviction","Copious","Core","Coupled","Courageous","Creative","Cuddle","Cultivate","Cure","Curious","Cute","Dazzling","Delight","Direct","Discover","Distinguished","Divine","Donate","Each Day","Eager","Earnest","Easy","Ecstasy","Effervescent","Efficient","Effortless","Electrifying","Elegance","Embrace","Encompassing","Encourage","Endorse","Energized","Energy","Enjoy","Enormous","Enthuse","Enthusiastic","Entirely","Essence","Established","Esteem","Everyday","Everyone","Excited","Exciting","Exhilarating","Expand","Explore","Express","Exquisite","Exultant","Faith","Familiar","Family","Famous","Feat","Fit","Flourish","Fortunate","Fortune","Freedom","Fresh","Friendship","Full","Funny","Gather","Generous","Genius","Genuine","Give","Glad","Glow","Good","Gorgeous","Grace","Graceful","Gratitude","Green","Grin","Group","Grow","Handsome","Happy","Harmony","Healed","Healing","Healthful","Healthy","Heart","Hearty","Heavenly","Helpful","Here","Highest Good","Hold","Holy","Honest","Honor","Hug","Idea","Ideal","Imaginative","Increase","Incredible","Independent","Ingenious","Innate","Innovate","Inspire","Instantaneous","Instinct","Intellectual","Intelligence","Intuitive","Inventive","Joined","Jovial","Joy","Jubilation","Keen","Key","Kind","Kiss","Knowledge","Laugh","Leader","Learn","Legendary","Let Go","Light","Lively","Love","Loveliness","Lucidity","Lucrative","Luminous","Maintain","Marvelous","Master","Meaningful","Meditate","Mend","Metamorphosis","Mind-Blowing","Miracle","Mission","Modify","Motivate","Moving","Natural","Nature","Nourish","Nourished","Novel","Now","Nurture","Nutritious","One","Open","Openhanded","Optimistic","Paradise","Party","Peace","Perfect","Phenomenon","Pleasure","Plenteous","Plentiful","Plenty","Plethora","Poise","Polish","Popular","Positive","Powerful","Prepared","Pretty","Principle","Productive","Project","Prominent","Prosperous","Protect","Proud","Purpose","Quest","Quick","Quiet","Ready","Recognize","Refinement","Refresh","Rejoice","Rejuvenate","Relax","Reliance","Rely","Remarkable","Renew","Renowned","Replenish","Resolution","Resound","Resources","Respect","Restore","Revere","Revolutionize","Rewarding","Rich","Robust","Rousing","Safe","Secure","See","Sensation","Serenity","Shift","Shine","Show","Silence","Simple","Sincerity","Smart","Smile","Smooth","Solution","Soul","Sparkling","Spirit","Spirited","Spiritual","Splendid","Spontaneous","Still","Stir","Strong","Style","Success","Sunny","Support","Sure","Surprise","Sustain","Synchronized","Team","Thankful","Therapeutic","Thorough","Thrilled","Thrive","Today","Together","Tranquil","Transform","Triumph","Trust","Truth","Unity","Unusual","Unwavering","Upbeat","Value","Vary","Venerate","Venture","Very","Vibrant","Victory","Vigorous","Vision","Visualize","Vital","Vivacious","Voyage","Wealthy","Welcome","Well","Whole","Wholesome","Willing","Wonder","Wonderful","Wondrous","Xanadu","Yes","Yippee","Young","Youth","Youthful","Zeal","Zest","Zing","Zip"];
    let randomIndex = this.getRandomIndexForArray(positiveWords);
    positiveWord = positiveWords[randomIndex];
    return positiveWord;
  }
  
  private getRandomBaseballTeams():string{
    let baseballTeam:string;
    let baseballTeams = ["Diamondbacks","Braves","Orioles","Red Sox","Cubs","White Sox","Reds","Indians","Rockies","Tigers","Astros","Royals","Angels","Dodgers","Marlins","Brewers","Twins","Mets","Yankees","Athletics","Phillies","Pirates","Padres","Giants","Mariners","Cardinals","Rays","Rangers","Blue Jays","Nationals"];
    let randomIndex = this.getRandomIndexForArray(baseballTeams);
    baseballTeam = baseballTeams[randomIndex];
    return baseballTeam;
  }

  private getRandomMultipleOf10():number{
    let max = 60/10;
    let randomMultipleOf10: number = (Math.floor(Math.random()*max)*10)+5;
    return randomMultipleOf10;
  }

  private getRandomIndexForArray(array:any[]):number{
    let randomIndex:number;
    randomIndex = Math.floor(Math.random() * (array.length-1))
    return randomIndex;
  }

  private getRandomNumberByMaxValue(max:number):number{
    let randomIndex:number;
    randomIndex = Math.floor(Math.random() * (max-1))
    return randomIndex;
  }
}
