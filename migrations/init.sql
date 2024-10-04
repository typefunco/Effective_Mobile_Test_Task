CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(16) UNIQUE,
    user_password VARCHAR,
    is_admin BOOLEAN
);

CREATE TABLE songs (
    song_id SERIAL PRIMARY KEY,
    song_name VARCHAR,
    song_text VARCHAR,
    release_date VARCHAR,
    song_link VARCHAR,
    song_author VARCHAR
);

INSERT INTO users (username, user_password, is_admin) VALUES
('Matvey', 'qwerty123', true),
('Vadim', 'qwerty123', true),
('lalala', 'qwerty123', true),
('string', 'qwerty123', false),
('Gamma', 'qwerty123', false),
('Beta', 'qwerty123', false),
('Alpha', 'qwerty123', false),
('Binance', 'qwerty123', false),
('Bybit', 'qwerty123', false),
('Crypto', 'qwerty123', false),
('Dex', 'qwerty123', false),
('Cex', 'qwerty123', false),
('Bitcoin', 'qwerty123', false),
('Etherium', 'qwerty123', false);

INSERT INTO songs (song_name, song_text, release_date, song_link, song_author) VALUES
('Updated Song Name', 'The waves crash against the shore|The sky is dark and foreboding|A ship sails lost in the tempest|Its sails torn and broken|Lightning splits the sky|The thunder roars aloud|Yet hope shines in the darkness|Guiding them to safety', '2024-10-04', 'https://example.com/stormy-ocean', 'Liam Turner'),
('City Lights', 'City lights flicker on|The streets alive with energy|Cars rush through the night|Neon signs paint the sky|People hurry home|Or out to meet friends|The city never sleeps|Always pulsing with life', '2021-11-15', 'https://example.com/city-lights', 'Mason Evans'),
('Lonely Path', 'The fire burns so bright|Its warmth fills the room|Crackling logs as they fall|Embers dance into the night|A story told by the flames|Whispers of love and longing|Yet it fades as dawn arrives|Leaving only the ashes behind', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett'),
('Burning Flame', 'The night is still and quiet|A whisper cuts through the dark|Secrets carried by the wind|Long forgotten and untold|Stars shine above like watchers|Their gaze fixed upon the earth|But no answer comes from above|Only silence remains', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett'),
('Silent Whisper', 'The mountain stands so tall|Its peak lost among the clouds|Echoes of a distant past|Ringing through the valleys below|Snow-covered slopes glisten bright|Bathed in the morning light|A lone traveler climbs the path|To see the world from above', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett'),
('Mountain Echo', 'Golden fields stretch so far|The harvest sways in the breeze|The sun paints the sky orange|As evening falls over the land|A farmer rests from his work|Gazing at the endless fields|His labor finally complete|And now he reaps what he sowed', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett'),
('Golden Fields', 'A journey without an end|The road winds ever onward|Mountains rise in the distance|Rivers flow beside the path|The sky changes with each step|Blue to gray and back again|A traveler with no destination|Seeking what cannot be found', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett'),
('Endless Journey', 'The garden blooms with color|Red petals fall from roses|Their fragrance fills the air|A gentle breeze carries them|They land softly on the grass|Like tears from a lover’s heart|Yet the beauty of the rose|Cannot be marred by sorrow', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett'),
('Crimson Petals', 'The river is frozen solid|Its surface like a mirror|Reflecting the winter sky|Clouds move across its face|A skater glides on the ice|Graceful as a bird in flight|Yet beneath the frozen shell|The water still flows below', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett'),
('Frozen River', 'Stars fill the sky above|A blanket of endless light|The moon rises slowly up|Bathing the world in silver|The night is peaceful and calm|Only the sound of the breeze|And the stars whisper softly|Their secrets to the earth below', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett'),
('Starry Night', 'I keep chasing after dreams|That seem just out of my reach|The harder I run, the further|They drift away from my grasp|But I will not stop running|Even if I never catch them|For the chase itself is life|And I live for the pursuit', '2020-06-19', 'https://example.com/lonely-path', 'Emma Bennett');

