// https://github.com/yckart/jquery-custom-animations
$.fn.blindLeftIn = function (duration, easing, complete) {
  return this.animate(
    { marginLeft: 0 },
    $.speed(duration, easing, complete)
  );
};

$.fn.blindLeftOut = function (duration, easing, complete) {
  return this.animate(
    { marginLeft: -this.outerWidth() },
    $.speed(duration, easing, complete)
  );
};

storage = false;
var fail, uid;
try {
  uid = new Date();
  (storage = window.localStorage).setItem(uid, uid);
  fail = storage.getItem(uid) != uid;
  storage.removeItem(uid);
  fail && (storage = false);
} catch(e) {}

if (!window.console) {
	console = { log: function() {} };
}

classInfos = [[768,"Rogue","Rogues",[150,100,10,0,15,15,15,10],[25,5,1,0,1.5,1.5,0.5,1],[720,252,50,25,75,75,40,50],[[0,"Classic",0],[834,"Bandit",14],[913,"Brigand",58],[916,"Eligible Bachelor",66]],1],[775,"Archer","Archers",[130,100,12,0,12,12,12,10],[25,5,1.5,0,1,1,0.5,1],[700,252,75,25,50,50,40,50],[[0,"Classic",1],[835,"Robin Hood",15],[851,"Little Helper",31],[855,"Cupid",35],[904,"Slime Archer",49],[888,"Ranger",56]],2],[782,"Wizard","Wizards",[100,100,12,0,10,15,12,12],[25,10,1.5,0,1,1.5,0.5,1],[670,385,75,25,50,75,40,60],[[0,"Classic",2],[836,"Merlin",16],[848,"Elder Wizard",28],[850,"Santa",30],[914,"Gentleman",55],[872,"Slime Wizard",46],[9012,"Witch",85]],3],[784,"Priest","Priests",[100,100,10,0,12,12,10,15],[25,10,1,0,1.5,1,0.5,1.5],[670,385,50,25,50,50,40,75],[[0,"Classic",3],[837,"Traditional",17],[849,"Robed Priest",29],[852,"Father Time",32],[905,"Slime Priest",45],[884,"Nun",52]],4],[797,"Warrior","Warriors",[200,100,15,0,7,10,10,10],[25,5,1.5,0,1,1,1.5,1],[770,252,75,25,50,50,75,50],[[0,"Classic",4],[838,"Juggernaut",18],[853,"Down Under",33],[883,"Shoveguy",51],[8967,"B.B. Wolf",72]],5],[798,"Knight","Knights",[200,100,15,0,7,10,10,10],[25,5,1.5,0,1,1,1.5,1],[770,252,50,40,50,50,75,50],[[0,"Classic",5],[839,"Knight of the Round",19],[903,"Slime Knight",44],[902,"Iceman",64]],6],[799,"Paladin","Paladins",[200,100,12,0,7,10,10,10],[25,5,1.5,0,1,1,0.5,1.5],[770,252,50,25,50,50,40,75],[[0,"Classic",6],[840,"Decorated Paladin",20],[854,"Founding Father",34],[915,"Bashing Bride",65],[885,"Holy Avenger",53]],7],[800,"Assassin","Assassins",[150,100,12,0,15,15,15,10],[25,5,1,0,1.5,1.5,0.5,1.5],[720,252,60,25,75,75,40,60],[[0,"Classic",7],[841,"Stealth",21],[912,"Agent",50]],8],[801,"Necromancer","Necromancers",[100,100,12,0,10,15,10,12],[25,10,1.5,0,1,1.5,0.5,1.5],[670,385,60,25,50,60,30,75],[[0,"Classic",8],[842,"Death Mage",22],[898,"Witch Doctor",60]],9],[802,"Huntress","Huntresses",[130,100,12,0,12,12,12,10],[25,5,1.5,0,1,1,0.5,1],[700,252,75,25,50,50,40,50],[[0,"Classic",9],[843,"Amazonian",23],[900,"Scarlett",62],[901,"Artemis",67],[8977,"Nexus no Miko",82]],10],[803,"Mystic","Mystics",[100,100,10,0,12,10,15,15],[25,10,1.5,0,1,1,0.5,1.5],[670,385,60,25,60,50,40,75],[[0,"Classic",10],[844,"Seer",24],[8968,"Lil Red",73]],11],[804,"Trickster","Tricksters",[150,100,10,0,12,15,12,12],[25,5,1.5,0,1.5,1.5,0.5,1],[720,252,65,25,75,75,40,60],[[0,"Classic",11],[845,"Loki",25],[917,"Deadly Vixen",63],[8979,"Drow Trickster",84],[8969,"King Knifeula",74]],12],[805,"Sorcerer","Sorcerers",[100,100,10,0,12,12,10,15],[25,10,1.5,0,1.5,1,1.5,1.5],[670,385,60,25,60,60,75,60],[[0,"Classic",12],[846,"Warlock",26],[899,"Sorceress",61],[8855,"Stanley the Spring Bunny",68]],13],[806,"Ninja","Ninjas",[150,100,15,0,10,12,10,12],[25,5,1.5,0,1,1.5,0.5,1.5],[720,252,70,25,60,70,40,70],[[0,"Classic",13],[847,"Shadow",27],[856,"Slime Ninja",36]],14]];
classInfoById = {};
for (var i = 0; i < classInfos.length; ++i) {
	classInfoById[classInfos[i][0]] = classInfos[i];
}

function statsTable(stats, bonuses, maxedArr, maxes, averages) {
  var
      table = ['<table class="stats-table">'],
      statNames = ["HP", "MP", "ATT", "DEF", "SPD", "VIT", "WIS", "DEX"];
  function addRow(from, to) {
    var allMaxed = true;
    table.push('<tr>');
    for(var i = from; i <= to; ++i) {
      addCell(statNames[i], maxedArr[i], stats[i], bonuses[i]);
      allMaxed &= maxedArr[i];
    }
    padRowToThreeCells(from, to);
    table.push('</tr>');
    if (!allMaxed) {
      addFromAvgRow(from, to);
      addToMaxRow(from, to);  
    }
  }
  function addFromAvgRow(from, to) {
    table.push('<tr>');
    for(var i = from; i <= to; ++i) {
      addFromAvgCell(maxedArr[i], stats[i], bonuses[i], averages[i]);
    }
    padRowToThreeCells(from, to);
    table.push('</tr>');
  }
  function addToMaxRow(from, to) {
    table.push('<tr>');
    for(var i = from; i <= to; ++i) {
      addToMaxCell(statNames[i], maxedArr[i], stats[i], bonuses[i], maxes[i]);
    }
    padRowToThreeCells(from, to);
    table.push('</tr>');
  }
  function addCell(name, maxed, stat, bonus) {
    table.push('<td');
    if (maxed) { table.push(' class="maxed"'); }
    table.push('>');
    table.push(name);
    table.push(': ');
    table.push(stat);
    if (bonus !== 0) {
      table.push('(');
      if (bonus > 0) {
        table.push('+');
      }
      table.push(bonus);
      table.push(')');
    }
    table.push('</td>');
  }
  function addFromAvgCell(maxed, stat, bonus, avg) {
    var fromAvg = stat - bonus - avg;
    table.push('<td class="from-avg' + ((fromAvg < 0) ? ' below-avg' : ' above-avg') + '">');
    if (!maxed) {
      if (fromAvg == 0) {
        table.push('average');
      } else {
        if (fromAvg > 0) {
          table.push('+');
        }
        table.push(fromAvg);
        table.push(' from avg');
      }
    }
    table.push('</td>');
  }
  function addToMaxCell(name, maxed, stat, bonus, max) {
    table.push('<td class="to-max">');
    if (!maxed) { 
      table.push(max - stat + bonus);
      if (name == 'HP' || name == 'MP') {
        table.push(' (');
        table.push(Math.ceil((max - stat + bonus) / 5));
        table.push(')');
      }
      table.push(' to max');
    }
    table.push('</td>');
  }
  function padRowToThreeCells(from, to) {
    for(i = 0; i < 3 - (to - from + 1); ++i) {
      table.push('<td></td>');
    }
  }
  addRow(0, 1);
  addRow(2, 4);
  addRow(5, 7);
  table.push('</table>');
  return table.join('');
}

function renderStats(tableId) {
  $("#" + tableId + " .player-stats").each(function() {
    var 
      self = $(this),
      mapping = [ 0, 1, 2, 3, 4, 6, 7, 5 ],
      stats = JSON.parse(self.attr("data-stats")),
      bonuses = JSON.parse(self.attr("data-bonuses")),
      playerClass = JSON.parse(self.attr("data-class")),
      level = JSON.parse(self.attr("data-level")),
      classInfo = classInfoById[playerClass],
      maxesOrig = classInfo[5],
      maxes = $.map(mapping, function(value) {
        return maxesOrig[value];
      }),
      maxed = [ false, false, false, false, false, false, false, false ];
      maxedCount = 0,
      gains = classInfo[4],
      starter = classInfo[3],
      average = $.map(gains, function(gain, i) {
        return starter[i] + gain * (level - 1);
      }),
      average = $.map(mapping, function(value) {
        return average[value];
      });
    $.each(maxes, function(index, max) {
      var base = stats[index] - bonuses[index];
      if (base == max) {
        maxed[index] = true;
        maxedCount += 1;
      }
    });
    self.html(maxedCount + '/8 <i class="icon icon-info-sign"></i>');
    self.popover({
      html: true,
      content: statsTable(stats, bonuses, maxed, maxes, average),
      trigger: "manual",
      template: '<div class="popover"><div class="arrow"></div><div class="popover-inner"><div class="popover-content"><p></p></div></div></div>'
    });
    self.parent().hover(function() { self.popover("show"); }, function () { self.popover("hide"); });
  });
}

function alternativesTable(alternatives) {
  var table = '<table class="alternatives-table"><thead><tr><th>Server</th><th>Price</th><th>Quantity</th><th>Time Left</th></tr></thead><tbody>';
  $.each(alternatives, function(index, row) {
    table += '<tr>';
    $.each(row, function(index, cell) {
      table += '<td>';
      if (index == 3) {
        if (cell != 2147483647) {
          table += '<';
          table += (cell + 1);
          table += ' min';
        } else {
          table += 'new';
        }
      } else {
        table += cell;
      }
      table += '</td>';
    });
    table += '</tr>';
  });
  table += '</tbody></table>';
  return table;
}

function renderAlternatives(tableId, columnIndex) {
  $("#" + tableId + " .cheapest-server").each(function() {
    var 
      self = $(this),
      alternatives = JSON.parse(self.attr("data-alternatives"));
    if (alternatives.length > 0) {
      self.append(' <i class="icon icon-info-sign"></i>');
      self.popover({
        html: true,
        content: alternativesTable(alternatives),
        trigger: "manual",
        title: 'Other servers',
        placement: 'left'
      });
      self.parent().hover(function() { self.popover("show"); }, function () { self.popover("hide"); });
    }
  });
}

function linkForItem(itemId) {
  var item = items[itemId] || items[-1];
  if ((item[1] == 0 || item[1] == 10 || item[1] == 26) && itemId != 3180) { // not equipment and not backpack
    return null; 
  }
  return '/wiki/' + item[0].toLowerCase().replace(/[\'\ ]/g, '-').replace(/\./g, '');
}

function renderItems(tableId) {
  makeItemsIn($("#" + tableId)).each(function() {
    $(this).wrap(function() { 
      var link = linkForItem($(this).attr('data-item'));
      if (link == null) { return false; }
      return '<a href="' + link + '" target="_blank"/>';
    });
  });
}

function makeItemsIn(element) {
  return $(".item", element).each(function() {
    makeItem($(this));
  });
}

function renderItemTable(item) {
  return [
    '<table class="stats-table"><tbody><tr><td>Fame Bonus:</td><td>',
    item[5],
    '%</td></tr><tr><td>Feed Power:</td><td>',
    item[6],
    '</td></tr></tbody></table>'].join('');
}

function makeItem(element) {
  var 
    item = items[element.attr("data-item")] || items[-1],
    tier = (item[1] == 10 || item[1] == 0 || item[1] == 26) ? '' : (item[2] == -1 ? ' UT' : ' T' + item[2]);
  element
    .css("background-position", "-" + item[3] + "px -" + item[4] + "px")
    .popover({
      title: item[0] + tier,
      trigger: 'hover',
      placement: 'top',
      html: true,
      content: renderItemTable(item) 
    });
}

function renderPets(tableId) {
  $("#" + tableId + " .pet").each(function() {
    var 
      self = $(this),
      item = items[self.attr("data-item")];
    if (!item) { return; }
    self
      .css("background-position", "-" + item[1] + "px -" + item[2] + "px")
      .attr("title", item[0]);
  });
}

function cssForColumn(tableId, columnIndex) {
  var column = $("#" + tableId + " td:nth-child(" + columnIndex + ")");
  for (var i = 2; i < arguments.length; i += 2) {
    column.css(arguments[i], arguments[i+1]);
  }
}

function renderArticle(tableId, columnIndex) {
  renderItems(tableId);
  $("#" + tableId + " td:nth-child(" + columnIndex + ") .item").each(function() { 
    var
      element = $(this);
      item = items[element.attr("data-item")] || items[-1],
      nextCell = element.closest('td').next();
    nextCell.text(item[0]);
  });
}

function alignColumnRight(tableId, columnIndex) {
  $("#" + tableId + " th:nth-child(" + columnIndex + ")")
    .css("text-align", "right");
  $("#" + tableId + " td:nth-child(" + columnIndex + ")")
    .css("text-align", "right");
}

function renderNumeric(tableId, columnIndex, withoutDiffs) {
  alignColumnRight(tableId, columnIndex);
	var 
		tds = $("#" + tableId + " td:nth-child(" + columnIndex + ")"),
		nodesToChange,
		children;
	nodesToChange = tds.map(function() {
		var node = $(this);
		while(node.children().length) {
			node = $((node.children())[0]);
		}
		return node;
	});
	nodesToChange.text(function(index, text) { 
		$(this).closest('td').data('value', text);
		return formatNumber(text); });
  if (!withoutDiffs) {
  	tds.hover(
	  	function() { showDiffs($(this)); },
      function() { hideDiffs($(this)); });
  }
}

function showDiffs(node) {
  var value = node.data('value');
  siblingsDo(node, 25, function(distance, otherValue) {
    if (otherValue == '') return;
    var 
      diff = otherValue - value,
      html = '<span class="diff"><span';
    if (diff < 0) { 
      html += ' class="diff-smaller">-';
    } else if(diff > 0) {
      html += ' class="diff-bigger">+';
    }
    if (diff % 1 != 0) { // if the difference is not an integer, round to 1 decimal place
      diff = diff.toFixed(1);
    }
    html += formatNumber(Math.abs(diff));
    html += '</span>';
    this.append(html);  
  });
}

function hideDiffs(node) {
  siblingsDo(node, 25, function() { 
    $('.diff', this).remove(); 
  });
}

function siblingsDo(node, count, callback) {
  siblingsSelectedDo(node, count, function(row) { return row.prev(); }, callback); 
  siblingsSelectedDo(node, count, function(row) { return row.next(); }, callback); 
}

function siblingsSelectedDo(node, count, selector, callback) {
  var 
    row = node.parent(),
    index = node.index();
  for(var i = 0; i < count; ++i) {
    row = selector(row);
    if (row.length) {
      var 
        td = row.find("td:nth-child(" + (index + 1) + ")"),
        value = td.data('value');
      callback.call(
        td,
        i + 1, // distance,
        value
      );
    } else {
      break;
    }
  }
}

function formatNumber(number) {
  var rgx = /(\d+)(\d{3})/;
  number += '';
  while (rgx.test(number)) {
    number = number.replace(rgx, '$1\u2009$2');
  }
  return number;
}

function formatTimeLeft(tableId, columnIndex) {
  $("#" + tableId + " td:nth-child(" + columnIndex + ")")
    .text(function(index, text) {
      var minutes = parseInt(text);
      if (minutes == 2147483647) {
        return 'new';
      } else if (minutes == 0) {
        return '<1 minute';
      } else {
        return '<' + (minutes + 1) + ' minutes';
      }
    }); 
}

function makeSortable(tableId, headers) {
  $("#" + tableId).tablesorter({
    headers: headers 
  });
}

function bookmarkPlayer(playerName) {
  bookmarkName(playerName, 'players');
}

function bookmarkGuild(guildName) {
  bookmarkName(guildName, 'guilds');
}

function bookmarkName(name, listName) {
  if (!storage) { return; }
  var list = storage[listName];
  if (list) {
    list = JSON.parse(list);
  } else {
    list = [];
  }
  var index = $.inArray(name, list);
  if (index != -1) {
    list.splice(index, 1);
  }
  list.unshift(name);
  list.splice(200, list.length - 200);
  storage[listName] = JSON.stringify(list);
}

function makeStars(tableId, columnIndex) {
  $("#" + tableId + " th:nth-child(" + columnIndex + ")")
    .css("text-align", "right");
  $("#" + tableId + " td:nth-child(" + columnIndex + ")")
    .css("text-align", "right")
    .each(function() {
      var 
        self = $(this),
        rank = parseInt(self.text());
        self.append('<div class="star star-' + colorForRank(rank) + '"></div>'); 
    });
}

function colorForRank(rank) {
  var color;
  $.each(
    [ 
      [ 13, 'light-blue' ], 
      [ 27, 'blue' ], 
      [ 41, 'red' ],
      [ 55, 'orange' ],
      [ 69, 'yellow' ],
      [ 70, 'white' ] 
    ], function(index, value) {
      color = value[1];
      if (rank <= value[0]) { return false; }
    });
  return color;
}

function renderDonationPopover(id, donations) {
  var 
    img = $('#' + id + ' img'),
    playerName = img.parent().prev().text(),
    srcCrownPlace = img.attr('src'),
    srcCrown = srcCrownPlace.slice(0, -1 * 'crown-place.png'.length) + 'crown.png';
  var title, content;
  if (donations.length) {
    function addDonation(index) {
      content +='<span class="timeago" title="' + donations[index] + '"></span>';
    }
    title = 'Thanks ' + playerName + '!';
    content = '<span>' + playerName + ' donated ';
    addDonation(0);
    for (var i = 1; i < donations.length - 1; ++i) {
      content += ', ';
      addDonation(i);
    }
    if (donations.length > 1) {
      content += ' and ';
      addDonation(donations.length - 1);
    }
    content +='</span>'
    content = $(content);
    $('.timeago', content).timeago();
  } else {
    title = playerName + " hasn't donated yet.";
    content = 'Click on the crown to donate';
    img.hover(
      function() { $(this).attr('src', srcCrown); },
      function() { $(this).attr('src', srcCrownPlace); }
    );
  }
  img
    .popover({
      html: true,
      title: title,
      content: content,
      trigger: 'hover'
    });
}

function addSearch(id, basePath, bookmarkKey) {
  var input = $("#" + id);
  input.change(function(evt) { 
      window.location.pathname = basePath + "/" + $(this).val(); 
    });
  if (storage) {
    var bookmarks = storage[bookmarkKey];
    if (bookmarks) {
      bookmarks = JSON.parse(storage[bookmarkKey]);
      input.typeahead({ 
        source: function(query) {
          return [ filterQuery(query) ].concat(bookmarks);
        }
      });
    }
  }
}

function addPlayerSearch(id) {
  addSearch(id, '/player', 'players');
}

function addGuildSearch(id) {
  addSearch(id, '/guild', 'guilds');
}

function initializeSearch(id) {
  var 
    $el = $('#' + id),
    $panel = $('.player-guild-toggle-panel', $el);
    $input = $('input', $el),
    $buttons = $('button', $el),
    $btnGroup = $('.btn-group', $el),
    $authPanel = $('.auth-panel');
  var shouldHide = true;
  var type = 'Player';
  var bookmarks = { 'Player': [], 'Guild': [] };

  if (storage) {
    bookmarks['Player'] = JSON.parse(storage['players'] || '[]');
    bookmarks['Guild'] = JSON.parse(storage['guilds'] || '[]');
  }

  $input
    .focus(function() {
      shouldHide = true;
      $el.addClass('input-append');
      $panel.show();
      $btnGroup.blindLeftIn('fast');
      $authPanel.fadeOut('fast');
      $input.attr('placeholder', type + ' search');
    })
    .blur(function() {
      if (shouldHide) {
        $btnGroup.blindLeftOut('fast', null, function() {
          $el.removeClass('input-append');
          $panel.hide();
        });
        $authPanel.fadeIn('fast');
        $input.attr('placeholder', 'Search');
      }
    })
    .typeahead({
      'source': function(query) {
        return [ filterQuery(query) ].concat(bookmarks[type]);
      },
      'updater': function(item) {
        search(item);
      }
    })
    .bind("paste", function() {
      setTimeout(function() {
        $input.autocomplete("search", $input.val());
      }, 0);
    });

  $buttons
    .mousedown(function() {
      shouldHide = false;
    })
    .click(function() {
      type = $(this).text();
      $input.trigger('focus');
    })
    .button();

  function search(item) {
    window.location.pathname = '/' + (type == 'Player' ? 'player' : 'guild') + '/' + item;
  }
}

function filterQuery(query) {
  return query.replace(/[^a-zA-Z ]/g, '').replace(/^ +| +$/g, '').replace(/ +/g, ' ');;
}

function indicateSelectedItem() {
  $("a[name=" + window.location.hash.slice(1) + "]").parent().css('border-left', '3px solid #777');
}

function makeChooserRedirector(id, prefix) {
  $("#" + id).change(function() {
    window.location.pathname = prefix + $(this).val();
  });
}

function escapeRegExp(str) {
  return str.replace(/[\-\[\]\/\{\}\(\)\*\+\?\.\\\^\$\|]/g, "\\$&");
}

function addAnchorsInDescription(id, customPattern) {
  var
    regex = new RegExp('(https?:\\/\\/)?(www\\.)?(youtube\\.com|youtu\\.be|imgur\\.com|i\\.imgur\\.com|puu\\.sh|twitch\\.tv|reddit\\.com|redd\\.it|github\\.com|community\\.kabam\\.com|forums\\.wildshadow\\.com|realmeye\\.com|bluenosersguide\\.weebly\\.com|pfiffel\\.com)(\\/[\\w\\/\\.\\?=&;_-]+)?' + (customPattern ? '|' + customPattern : ''), 'gi'),
    div = $("#" + id);
  $('.description-line', div).each(function() {
    $(this).html($(this).html().replace(regex, function(match) {
      var href = match;
      if (!/^https?:\/\//.test(match)) {
        href = 'http://' + match;
      }
      var extra = new RegExp(escapeRegExp(window.location.host)).test(match) ? ' rel="nofollow"' : ' rel="nofollow external"';
      return '<a href="' + encodeURI(href) + '"' + extra + '>' + match + '</a>';
    }))
  });
}

function renderMail(id) {
  $("#" + id).attr(
      "href", 
      $.map(
        [ 109,97,105,108,116,111,58,105,110,102,111,64,114,101,97,108,109,101,121,101,46,99,111,109 ], 
        function(each) { 
          return String.fromCharCode(each); 
        }).join("")
    );
}

function initializeLoginButton(id) {
  $('#' + id).click(function() {
    $(this).button('loading');
  });
}

function initializeLogin(id, loginSpec) {
  var 
    $modal = $('#' + id),
    $loginButton = $('button', $modal),
    submitted = false,
    https = window.location.protocol == 'https:',
    login = window.location.hash == '#login';
  $form = $('form', $modal);
  $loginButton.click(function() {
    $form.trigger('submit');
  });
  $("input").keypress(function(evt) {
    if (evt.which == 13) {
      evt.preventDefault();
      $form.trigger('submit');
    }
  });
  $form.submit(function(evt) {
    $loginButton.button('loading');
    if (submitted) {
      return true;
    }
    var 
      username = $(':input[name=username]', $form).val(),
      password = $(':input[name=password]', $form).val(),
      bindToIp = $(':input[name=bindToIp]', $form).is(":checked") ? 't' : 'f';
    if (!username || !password) {
      return false;
    }
    callSpec(loginSpec, {
        username: username,
        password: password,
        bindToIp: bindToIp
    }).done(function(data, status, xhr) {
      if (data == "OK") {
        submitted = true;
        $form.submit();
      } else {
        $loginButton.button('reset');
        $('.modal-footer .alert', $modal)
          .text('Invalid username or password!')
          .show();
      };
    }).fail(function() {
      $('.modal-footer .alert', $modal)
        .text('An error occured. Please try again!')
        .show();
    });
    evt.preventDefault();
  });
}

function initializeLogout(id) {
  $('#' + id).click(function(evt) {
    var href = window.location.pathname;
    $.ajax({
      type: 'POST',
      url: '/logout'
    }).done(function(data) {
      var hashMarkIndex = href.indexOf('#');
      if (hashMarkIndex != -1) {
        href = href.slice(0, hashMarkIndex);
      }
      window.location.href = href;
    });
  });
}

$.tablesorter.addParser({
  id: 'guildRank',
  is: function (s) { return false; },
  format: function (s) {
    switch (s) {
      case 'Founder': return 5;
      case 'Leader': return 4;
      case 'Officer': return 3;
      case 'Member': return 2;
      case 'Initiate': return 1;
      default: return 0;
    }
  },
  type: 'numeric'
});

$.tablesorter.addParser({
  id: 'petRarity',
  is: function (s) { return false; },
  format: function (s) {
    switch (s) {
      case 'Divine': return 5;
      case 'Legendary': return 4;
      case 'Rare': return 3;
      case 'Uncommon': return 2;
      case 'Common': return 1;
      default: return 0;
    }
  },
  type: 'numeric'
});

$.tablesorter.addParser({
  id: 'accountCreated',
  is: function (s) { return false; },
  format: function (s, table, td) {
    return $('span', td).data('rank');
  },
  type: 'numeric'
});

$.tablesorter.addParser({
  id: 'guildStarDistribution',
  is: function (s) { return false; },
  format: function (s, table, td) {
    return $('.guild-star-distribution', td).data('sorter');
  },
  type: 'numeric'
});

$(function() {
  $("span.numeric").text(function(index, text) { 
    return formatNumber(text); 
  });
  $("span.timeago").timeago();
  $('.dropdown-menu').on('touchstart.dropdown.data-api', function(e) { e.stopPropagation(); });
});

function initializeEditDescription(id, spec) {
  var
    button = $("#" + id + " > button"),
    editor = $("#" + id + " .modal");
  button.click(function() {
    editor.modal("show");
  });
  $('button', editor).click(function() {
    var 
      data = {};
    $.each(['line1', 'line2', 'line3'], function() { 
      data[this] = $('input[name=' + this + ']', editor).val();
    });
    callSpecAndReload(spec, data);
  });
}

function initializeClickHandlerWithAction(id, spec) { 
  $('#' + id).click(function() {
    callSpecAndReload(spec);
  });
}

function markBadMerchants(id, columnIndex) {
  $('#' + id + ' tbody tr:has(a.bad-merchant)').css('opacity', '0.2');
}

function petAbilityTable(level, maxLevel, points, pointsToNextLevel, pointsToMaxLevel) {
  var html = ['<table class="stats-table pet-ability-stats-table">'];
  addRow('Level:', level);
  addRow('Points:', points);
  if (level != maxLevel) {
    addRow('<abbr title="Feed power needed for Next Level">Next</abbr>:', pointsToNextLevel + ' fp');
    addRow('<abbr title="Feed power needed for Max Level(' + maxLevel + ')">Max</abbr>:', pointsToMaxLevel + ' fp');
  }
  html.push('</table>');
  return html.join('');

  function addRow(label, value) {
    html.push('<tr><td>');
    html.push(label);
    html.push('</td><td>');
    html.push(formatNumber(value));
    html.push('</td></tr>');
  }
}

function renderPetAbilityPopover(tableId, columnIndex, abilityIndex) {
  $("#" + tableId + " td:nth-child(" + columnIndex + ") .pet-ability").each(function() {
    var 
      self = $(this),
      data = self.data('pet-ability');
    self.html(data[1] + ' <i class="icon icon-info-sign"></i>');
    self.popover({
      title: data[0],
      html: true,
      content: petAbilityTable(data[1], data[2], data[3], data[4], data[5]),
      trigger: "click",
    });
  });
}

function completeData(data) {
  var session = document.cookie.match(/session=([0-9a-zA-Z]+)/);
  if (session && session[1]) {
    data.session = session[1];
  }
  return data;
}

function callSpec(spec, data) {
  $.extend(spec.data, data);
  completeData(spec.data);
  return $.ajax(spec);
}

function callSpecAndReload(spec, data, callback) {
  var href = window.location.href;
  callSpec(spec, data)
    .done(function(data) {
      if (!callback || callback(data)) {
        var hashMarkIndex = href.indexOf('#');
        if (hashMarkIndex != -1) {
          href = href.slice(0, hashMarkIndex);
        }
        window.location.href = href;
      }
    });
}

RealmEye = {
  converter: function() {
    return this.converterObject ||
      (this.converterObject = new Showdown.converter());
  },
  sanitizer: function() {
    return this.sanitizerObject ||
      (this.sanitizerObject = new Sanitize(Sanitize.Config.RELAXED));
  },
  sanitize: function(node, text) {
    var 
      converted = this.converter().makeHtml(text || node.text()),
      div = $('<div>' + converted + '</div>');
    var sanitized = this.sanitizer().clean_node(div[0]);
    node
      .empty()
      .append(sanitized); 
  },
  initializeEditor: function(textarea, button, checkbox, preview, ajaxSpec, additionalDataCallback) {
    if (textarea.data('ajaxSpec') === ajaxSpec) {
      return;
    }
    textarea.data('ajaxSpec', ajaxSpec);
    textarea.off();
    button.off();
    checkbox.off();
    if (ajaxSpec.data.body) {
      textarea.val(ajaxSpec.data.body);
    }
    checkbox.prop('checked', ajaxSpec.data.markdown);
    checkbox.change(updatePreview);
    textarea.on('input propertychange', updatePreview); 
    button.click(function() {
      button
        .attr('disabled', 'disabled')
        .addClass('disabled')
        .off();
      var data = { 
        body: textarea.val(),
        markdown: (checkbox.is(':checked') ? 't' : 'f')
      };
      if (additionalDataCallback) {
        additionalDataCallback(data);
      }
      callSpecAndReload(ajaxSpec, data);
    });
    updatePreview();

    function updatePreview() {
      if (checkbox.is(':checked')) {
        preview.addClass('markdown');
        RealmEye.sanitize(preview, textarea.val());
      } else {
        preview.removeClass('markdown');
        preview.text(textarea.val());
      }
    }
  }
}

function initializeHideCookieBanner(id) {
  $("#" + id + " .close").click(function () {
    document.cookie = 'hideCookieBanner=true;path=/;expires=Wed, 01 Jan 2020 00:00:00 GMT';
    $("#" + id).hide();
  });
}

function convertColumnToLocalTime(tableId, columnIndex) {
  $("#" + tableId + " td:nth-child(" + columnIndex + ")").text(function(index, text) {
    if (!text) { return ''; }
    var date = new Date(text);
    return ( 
      date.getFullYear() + '-' +
      padWithZeros(date.getMonth() + 1, 2) + '-' +
      padWithZeros(date.getDate(), 2) + ' ' +
      padWithZeros(date.getHours(), 2) + ':' +
      padWithZeros(date.getMinutes(), 2));
  });
}

function padWithZeros(n, width) {
  n = n + '';
  return n.length >= width ? n : new Array(width - n.length + 1).join('0') + n;
}

function initializeAlertCloseButton(id, alertVersion) {
  $('#' + id).click(function() {
    document.cookie = 'closedAlertVersion=' + alertVersion + ';path=/;expires=Wed, 01 Jan 2020 00:00:00 GMT';
  });
}
