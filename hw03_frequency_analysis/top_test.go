package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10WithAsteriskIsCompleted(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10("", true), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"а",         // 8
			"он",        // 8
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"в",         // 4
			"его",       // 4
			"если",      // 4
			"кристофер", // 4
			"не",        // 4
		}
		require.Equal(t, expected, Top10(text, true))
	})
}

func TestTop10WithAsteriskIsUncompleted(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10("", false), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"он",        // 8
			"а",         // 6
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"-",         // 4
			"Кристофер", // 4
			"если",      // 4
			"не",        // 4
			"то",        // 4
		}
		require.Equal(t, expected, Top10(text, false))
	})
}

func TestClearAndSplitWhenTaskWithAsteriskIsUncompleted(t *testing.T) {
	input := "На   \tдворе -  трава, на  траве - дрова. Не руби \n дрова на траве \nдвора!"
	expected := []string{
		"На", "дворе", "-", "трава,", "на", "траве", "-", "дрова.",
		"Не", "руби", "дрова", "на", "траве", "двора!",
	}
	result := ClearAndSplit(input, false)

	require.Len(t, result, 14)
	require.Equal(t, result, expected)
}

func TestClearAndSplitWhenTaskWithAsteriskIsCompleted(t *testing.T) {
	input := "На   \tдворе -  трава, на  траве - дрова. Не руби \n дрова на траве \nдвора!"
	expected := []string{
		"на", "дворе", "трава", "на", "траве", "дрова",
		"не", "руби", "дрова", "на", "траве", "двора",
	}
	result := ClearAndSplit(input, true)

	require.Len(t, result, 12)
	require.Equal(t, result, expected)
}

func TestCountWordsAndRepetition(t *testing.T) {
	input := []string{
		"на", "дворе", "трава", "на", "траве", "дрова",
		"не", "руби", "дрова", "на", "траве", "двора",
	}
	expected := []dimension{
		{word: "на", count: 3},
		{word: "дворе", count: 1},
		{word: "трава", count: 1},
		{word: "траве", count: 2},
		{word: "дрова", count: 2},
		{word: "не", count: 1},
		{word: "руби", count: 1},
		{word: "двора", count: 1},
	}

	result := countWordsAndRepetition(input)
	require.Len(t, result, len(expected))
}

func TestGet10TopCountedWordsWithSmallerWords(t *testing.T) {
	input := []dimension{
		{word: "на", count: 3},
		{word: "дворе", count: 1},
		{word: "трава", count: 1},
		{word: "траве", count: 2},
		{word: "дрова", count: 2},
		{word: "не", count: 1},
		{word: "руби", count: 1},
		{word: "двора", count: 1},
	}
	expected := []string{
		"на", "дрова", "траве", "двора", "дворе", "не", "руби", "трава",
	}
	result := get10TopCountedWords(input)

	require.Len(t, result, len(expected))
	require.Equal(t, result, expected)
}
