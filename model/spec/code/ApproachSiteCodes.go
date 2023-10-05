package code

import (
	"fmt"
	"strings"
)

/* TODO change type to string */
type ApproachSiteCodes int

const (
	ApproachSiteCodes000 ApproachSiteCodes = iota
	ApproachSiteCodes001
	ApproachSiteCodes002
	ApproachSiteCodes003
	ApproachSiteCodes004
	ApproachSiteCodes005
	ApproachSiteCodes006
	ApproachSiteCodes007
	ApproachSiteCodes008
	ApproachSiteCodes009
	ApproachSiteCodes010
	ApproachSiteCodes011
	ApproachSiteCodes012
	ApproachSiteCodes013
	ApproachSiteCodes014
	ApproachSiteCodes015
	ApproachSiteCodes016
	ApproachSiteCodes017
	ApproachSiteCodes018
	ApproachSiteCodes019
	ApproachSiteCodes020
	ApproachSiteCodes021
	ApproachSiteCodes022
	ApproachSiteCodes023
	ApproachSiteCodes024
	ApproachSiteCodes025
	ApproachSiteCodes026
	ApproachSiteCodes027
	ApproachSiteCodes028
	ApproachSiteCodes029
	ApproachSiteCodes030
	ApproachSiteCodes031
	ApproachSiteCodes032
	ApproachSiteCodes033
	ApproachSiteCodes034
	ApproachSiteCodes035
	ApproachSiteCodes036
	ApproachSiteCodes037
	ApproachSiteCodes038
	ApproachSiteCodes039
	ApproachSiteCodes040
	ApproachSiteCodes041
	ApproachSiteCodes042
	ApproachSiteCodes043
	ApproachSiteCodes044
	ApproachSiteCodes045
	ApproachSiteCodes046
	ApproachSiteCodes047
	ApproachSiteCodes048
	ApproachSiteCodes049
	ApproachSiteCodes050
	ApproachSiteCodes051
	ApproachSiteCodes052
	ApproachSiteCodes053
	ApproachSiteCodes054
	ApproachSiteCodes055
	ApproachSiteCodes056
	ApproachSiteCodes057
	ApproachSiteCodes058
	ApproachSiteCodes059
	ApproachSiteCodes060
	ApproachSiteCodes061
	ApproachSiteCodes062
	ApproachSiteCodes063
	ApproachSiteCodes064
	ApproachSiteCodes065
	ApproachSiteCodes066
	ApproachSiteCodes067
	ApproachSiteCodes068
	ApproachSiteCodes069
	ApproachSiteCodes070
	ApproachSiteCodes071
	ApproachSiteCodes072
	ApproachSiteCodes073
	ApproachSiteCodes074
	ApproachSiteCodes075
	ApproachSiteCodes076
	ApproachSiteCodes077
	ApproachSiteCodes078
	ApproachSiteCodes079
	ApproachSiteCodes080
	ApproachSiteCodes081
	ApproachSiteCodes082
	ApproachSiteCodes083
	ApproachSiteCodes084
	ApproachSiteCodes085
	ApproachSiteCodes086
	ApproachSiteCodes087
	ApproachSiteCodes088
	ApproachSiteCodes089
	ApproachSiteCodes090
	ApproachSiteCodes091
	ApproachSiteCodes092
	ApproachSiteCodes093
	ApproachSiteCodes094
	ApproachSiteCodes095
	ApproachSiteCodes096
	ApproachSiteCodes097
	ApproachSiteCodes098
	ApproachSiteCodes099
	ApproachSiteCodes100
	ApproachSiteCodes101
	ApproachSiteCodes102
	ApproachSiteCodes103
	ApproachSiteCodes104
	ApproachSiteCodes105
	ApproachSiteCodes106
	ApproachSiteCodes107
	ApproachSiteCodes108
	ApproachSiteCodes109
	ApproachSiteCodes110
	ApproachSiteCodes111
	ApproachSiteCodes112
	ApproachSiteCodes113
	ApproachSiteCodes114
	ApproachSiteCodes115
	ApproachSiteCodes116
	ApproachSiteCodes117
	ApproachSiteCodes118
	ApproachSiteCodes119
	ApproachSiteCodes120
	ApproachSiteCodes121
	ApproachSiteCodes122
	ApproachSiteCodes123
	ApproachSiteCodes124
	ApproachSiteCodes125
	ApproachSiteCodes126
	ApproachSiteCodes127
	ApproachSiteCodes128
	ApproachSiteCodes129
	ApproachSiteCodes130
	ApproachSiteCodes131
	ApproachSiteCodes132
	ApproachSiteCodes133
	ApproachSiteCodes134
	ApproachSiteCodes135
	ApproachSiteCodes136
	ApproachSiteCodes137
	ApproachSiteCodes138
	ApproachSiteCodes139
	ApproachSiteCodes140
	ApproachSiteCodes141
	ApproachSiteCodes142
	ApproachSiteCodes143
	ApproachSiteCodes144
	ApproachSiteCodes145
	ApproachSiteCodes146
	ApproachSiteCodes147
	ApproachSiteCodes148
	ApproachSiteCodes149
	ApproachSiteCodes150
	ApproachSiteCodes151
	ApproachSiteCodes152
	ApproachSiteCodes153
	ApproachSiteCodes154
	ApproachSiteCodes155
	ApproachSiteCodes156
	ApproachSiteCodes157
	ApproachSiteCodes158
	ApproachSiteCodes159
	ApproachSiteCodes160
	ApproachSiteCodes161
	ApproachSiteCodes162
	ApproachSiteCodes163
	ApproachSiteCodes164
	ApproachSiteCodes165
	ApproachSiteCodes166
	ApproachSiteCodes167
	ApproachSiteCodes168
	ApproachSiteCodes169
	ApproachSiteCodes170
	ApproachSiteCodes171
	ApproachSiteCodes172
	ApproachSiteCodes173
	ApproachSiteCodes174
	ApproachSiteCodes175
	ApproachSiteCodes176
	ApproachSiteCodes177
	ApproachSiteCodes178
	ApproachSiteCodes179
	ApproachSiteCodes180
	ApproachSiteCodes181
	ApproachSiteCodes182
	ApproachSiteCodes183
	ApproachSiteCodes184
	ApproachSiteCodes185
	ApproachSiteCodes186
	ApproachSiteCodes187
	ApproachSiteCodes188
	ApproachSiteCodes189
	ApproachSiteCodes190
	ApproachSiteCodes191
	ApproachSiteCodes192
	ApproachSiteCodes193
	ApproachSiteCodes194
	ApproachSiteCodes195
	ApproachSiteCodes196
	ApproachSiteCodes197
	ApproachSiteCodes198
	ApproachSiteCodes199
	ApproachSiteCodes200
	ApproachSiteCodes201
	ApproachSiteCodes202
	ApproachSiteCodes203
	ApproachSiteCodes204
	ApproachSiteCodes205
	ApproachSiteCodes206
	ApproachSiteCodes207
	ApproachSiteCodes208
	ApproachSiteCodes209
	ApproachSiteCodes210
	ApproachSiteCodes211
	ApproachSiteCodes212
	ApproachSiteCodes213
	ApproachSiteCodes214
	ApproachSiteCodes215
	ApproachSiteCodes216
	ApproachSiteCodes217
	ApproachSiteCodes218
	ApproachSiteCodes219
	ApproachSiteCodes220
	ApproachSiteCodes221
	ApproachSiteCodes222
	ApproachSiteCodes223
	ApproachSiteCodes224
	ApproachSiteCodes225
	ApproachSiteCodes226
	ApproachSiteCodes227
	ApproachSiteCodes228
	ApproachSiteCodes229
	ApproachSiteCodes230
	ApproachSiteCodes231
	ApproachSiteCodes232
	ApproachSiteCodes233
	ApproachSiteCodes234
	ApproachSiteCodes235
	ApproachSiteCodes236
	ApproachSiteCodes237
	ApproachSiteCodes238
	ApproachSiteCodes239
	ApproachSiteCodes240
	ApproachSiteCodes241
	ApproachSiteCodes242
	ApproachSiteCodes243
	ApproachSiteCodes244
	ApproachSiteCodes245
	ApproachSiteCodes246
	ApproachSiteCodes247
	ApproachSiteCodes248
	ApproachSiteCodes249
	ApproachSiteCodes250
	ApproachSiteCodes251
	ApproachSiteCodes252
	ApproachSiteCodes253
	ApproachSiteCodes254
	ApproachSiteCodes255
	ApproachSiteCodes256
	ApproachSiteCodes257
	ApproachSiteCodes258
	ApproachSiteCodes259
	ApproachSiteCodes260
	ApproachSiteCodes261
	ApproachSiteCodes262
	ApproachSiteCodes263
	ApproachSiteCodes264
	ApproachSiteCodes265
	ApproachSiteCodes266
	ApproachSiteCodes267
	ApproachSiteCodes268
	ApproachSiteCodes269
	ApproachSiteCodes270
	ApproachSiteCodes271
	ApproachSiteCodes272
	ApproachSiteCodes273
	ApproachSiteCodes274
	ApproachSiteCodes275
	ApproachSiteCodes276
	ApproachSiteCodes277
	ApproachSiteCodes278
	ApproachSiteCodes279
	ApproachSiteCodes280
	ApproachSiteCodes281
	ApproachSiteCodes282
	ApproachSiteCodes283
	ApproachSiteCodes284
	ApproachSiteCodes285
	ApproachSiteCodes286
	ApproachSiteCodes287
	ApproachSiteCodes288
	ApproachSiteCodes289
	ApproachSiteCodes290
	ApproachSiteCodes291
	ApproachSiteCodes292
	ApproachSiteCodes293
	ApproachSiteCodes294
	ApproachSiteCodes295
	ApproachSiteCodes296
	ApproachSiteCodes297
	ApproachSiteCodes298
	ApproachSiteCodes299
	ApproachSiteCodes300
	ApproachSiteCodes301
	ApproachSiteCodes302
	ApproachSiteCodes303
	ApproachSiteCodes304
	ApproachSiteCodes305
	ApproachSiteCodes306
	ApproachSiteCodes307
	ApproachSiteCodes308
	ApproachSiteCodes309
	ApproachSiteCodes310
	ApproachSiteCodes311
	ApproachSiteCodes312
	ApproachSiteCodes313
	ApproachSiteCodes314
	ApproachSiteCodes315
	ApproachSiteCodes316
	ApproachSiteCodes317
	ApproachSiteCodes318
	ApproachSiteCodes319
	ApproachSiteCodes320
	ApproachSiteCodes321
	ApproachSiteCodes322
	ApproachSiteCodes323
	ApproachSiteCodes324
	ApproachSiteCodes325
	ApproachSiteCodes326
	ApproachSiteCodes327
	ApproachSiteCodes328
	ApproachSiteCodes329
	ApproachSiteCodes330
	ApproachSiteCodes331
	ApproachSiteCodes332
	ApproachSiteCodes333
	ApproachSiteCodes334
	ApproachSiteCodes335
	ApproachSiteCodes336
	ApproachSiteCodes337
	ApproachSiteCodes338
	ApproachSiteCodes339
	ApproachSiteCodes340
	ApproachSiteCodes341
	ApproachSiteCodes342
	ApproachSiteCodes343
	ApproachSiteCodes344
	ApproachSiteCodes345
	ApproachSiteCodes346
	ApproachSiteCodes347
	ApproachSiteCodes348
	ApproachSiteCodes349
	ApproachSiteCodes350
	ApproachSiteCodes351
	ApproachSiteCodes352
	ApproachSiteCodes353
	ApproachSiteCodes354
	ApproachSiteCodes355
	ApproachSiteCodes356
	ApproachSiteCodes357
	ApproachSiteCodes358
	ApproachSiteCodes359
	ApproachSiteCodes360
	ApproachSiteCodes361
	ApproachSiteCodes362
	ApproachSiteCodes363
	ApproachSiteCodes364
	ApproachSiteCodes365
	ApproachSiteCodes366
	ApproachSiteCodes367
	ApproachSiteCodes368
	ApproachSiteCodes369
	ApproachSiteCodes370
	ApproachSiteCodes371
	ApproachSiteCodes372
	ApproachSiteCodes373
	ApproachSiteCodes374
	ApproachSiteCodes375
	ApproachSiteCodes376
	ApproachSiteCodes377
	ApproachSiteCodes378
	ApproachSiteCodes379
	ApproachSiteCodes380
	ApproachSiteCodes381
	ApproachSiteCodes382
	ApproachSiteCodes383
	ApproachSiteCodes384
	ApproachSiteCodes385
	ApproachSiteCodes386
	ApproachSiteCodes387
	ApproachSiteCodes388
	ApproachSiteCodes389
	ApproachSiteCodes390
	ApproachSiteCodes391
	ApproachSiteCodes392
	ApproachSiteCodes393
	ApproachSiteCodes394
	ApproachSiteCodes395
	ApproachSiteCodes396
	ApproachSiteCodes397
	ApproachSiteCodes398
	ApproachSiteCodes399
	ApproachSiteCodes400
	ApproachSiteCodes401
	ApproachSiteCodes402
	ApproachSiteCodes403
	ApproachSiteCodes404
	ApproachSiteCodes405
	ApproachSiteCodes406
	ApproachSiteCodes407
	ApproachSiteCodes408
	ApproachSiteCodes409
	ApproachSiteCodes410
	ApproachSiteCodes411
	ApproachSiteCodes412
	ApproachSiteCodes413
	ApproachSiteCodes414
	ApproachSiteCodes415
	ApproachSiteCodes416
	ApproachSiteCodes417
	ApproachSiteCodes418
	ApproachSiteCodes419
	ApproachSiteCodes420
	ApproachSiteCodes421
	ApproachSiteCodes422
	ApproachSiteCodes423
	ApproachSiteCodes424
	ApproachSiteCodes425
	ApproachSiteCodes426
	ApproachSiteCodes427
	ApproachSiteCodes428
	ApproachSiteCodes429
	ApproachSiteCodes430
	ApproachSiteCodes431
	ApproachSiteCodes432
	ApproachSiteCodes433
	ApproachSiteCodes434
	ApproachSiteCodes435
	ApproachSiteCodes436
	ApproachSiteCodes437
	ApproachSiteCodes438
	ApproachSiteCodes439
	ApproachSiteCodes440
	ApproachSiteCodes441
	ApproachSiteCodes442
	ApproachSiteCodes443
	ApproachSiteCodes444
	ApproachSiteCodes445
	ApproachSiteCodes446
	ApproachSiteCodes447
	ApproachSiteCodes448
	ApproachSiteCodes449
	ApproachSiteCodes450
	ApproachSiteCodes451
	ApproachSiteCodes452
	ApproachSiteCodes453
	ApproachSiteCodes454
	ApproachSiteCodes455
	ApproachSiteCodes456
	ApproachSiteCodes457
	ApproachSiteCodes458
	ApproachSiteCodes459
	ApproachSiteCodes460
	ApproachSiteCodes461
	ApproachSiteCodes462
	ApproachSiteCodes463
	ApproachSiteCodes464
	ApproachSiteCodes465
	ApproachSiteCodes466
	ApproachSiteCodes467
	ApproachSiteCodes468
	ApproachSiteCodes469
	ApproachSiteCodes470
	ApproachSiteCodes471
	ApproachSiteCodes472
	ApproachSiteCodes473
	ApproachSiteCodes474
	ApproachSiteCodes475
	ApproachSiteCodes476
	ApproachSiteCodes477
	ApproachSiteCodes478
	ApproachSiteCodes479
	ApproachSiteCodes480
	ApproachSiteCodes481
	ApproachSiteCodes482
	ApproachSiteCodes483
	ApproachSiteCodes484
	ApproachSiteCodes485
	ApproachSiteCodes486
	ApproachSiteCodes487
	ApproachSiteCodes488
	ApproachSiteCodes489
	ApproachSiteCodes490
	ApproachSiteCodes491
	ApproachSiteCodes492
	ApproachSiteCodes493
	ApproachSiteCodes494
	ApproachSiteCodes495
	ApproachSiteCodes496
	ApproachSiteCodes497
	ApproachSiteCodes498
	ApproachSiteCodes499
	ApproachSiteCodes500
	ApproachSiteCodes501
	ApproachSiteCodes502
	ApproachSiteCodes503
	ApproachSiteCodes504
	ApproachSiteCodes505
	ApproachSiteCodes506
	ApproachSiteCodes507
	ApproachSiteCodes508
	ApproachSiteCodes509
	ApproachSiteCodes510
	ApproachSiteCodes511
	ApproachSiteCodes512
	ApproachSiteCodes513
	ApproachSiteCodes514
	ApproachSiteCodes515
	ApproachSiteCodes516
	ApproachSiteCodes517
	ApproachSiteCodes518
	ApproachSiteCodes519
	ApproachSiteCodes520
	ApproachSiteCodes521
	ApproachSiteCodes522
	ApproachSiteCodes523
	ApproachSiteCodes524
	ApproachSiteCodes525
	ApproachSiteCodes526
	ApproachSiteCodes527
	ApproachSiteCodes528
	ApproachSiteCodes529
	ApproachSiteCodes530
	ApproachSiteCodes531
	ApproachSiteCodes532
	ApproachSiteCodes533
	ApproachSiteCodes534
	ApproachSiteCodes535
	ApproachSiteCodes536
	ApproachSiteCodes537
	ApproachSiteCodes538
	ApproachSiteCodes539
	ApproachSiteCodes540
	ApproachSiteCodes541
	ApproachSiteCodes542
	ApproachSiteCodes543
	ApproachSiteCodes544
	ApproachSiteCodes545
	ApproachSiteCodes546
	ApproachSiteCodes547
	ApproachSiteCodes548
	ApproachSiteCodes549
	ApproachSiteCodes550
	ApproachSiteCodes551
	ApproachSiteCodes552
	ApproachSiteCodes553
	ApproachSiteCodes554
	ApproachSiteCodes555
	ApproachSiteCodes556
	ApproachSiteCodes557
	ApproachSiteCodes558
	ApproachSiteCodes559
	ApproachSiteCodes560
	ApproachSiteCodes561
	ApproachSiteCodes562
	ApproachSiteCodes563
	ApproachSiteCodes564
	ApproachSiteCodes565
	ApproachSiteCodes566
	ApproachSiteCodes567
	ApproachSiteCodes568
	ApproachSiteCodes569
	ApproachSiteCodes570
	ApproachSiteCodes571
	ApproachSiteCodes572
	ApproachSiteCodes573
	ApproachSiteCodes574
	ApproachSiteCodes575
	ApproachSiteCodes576
	ApproachSiteCodes577
	ApproachSiteCodes578
	ApproachSiteCodes579
	ApproachSiteCodes580
	ApproachSiteCodes581
	ApproachSiteCodes582
	ApproachSiteCodes583
	ApproachSiteCodes584
	ApproachSiteCodes585
	ApproachSiteCodes586
	ApproachSiteCodes587
	ApproachSiteCodes588
	ApproachSiteCodes589
	ApproachSiteCodes590
	ApproachSiteCodes591
	ApproachSiteCodes592
	ApproachSiteCodes593
	ApproachSiteCodes594
	ApproachSiteCodes595
	ApproachSiteCodes596
	ApproachSiteCodes597
	ApproachSiteCodes598
	ApproachSiteCodes599
	ApproachSiteCodes600
	ApproachSiteCodes601
	ApproachSiteCodes602
	ApproachSiteCodes603
	ApproachSiteCodes604
	ApproachSiteCodes605
	ApproachSiteCodes606
	ApproachSiteCodes607
	ApproachSiteCodes608
	ApproachSiteCodes609
	ApproachSiteCodes610
	ApproachSiteCodes611
	ApproachSiteCodes612
	ApproachSiteCodes613
	ApproachSiteCodes614
	ApproachSiteCodes615
	ApproachSiteCodes616
	ApproachSiteCodes617
	ApproachSiteCodes618
	ApproachSiteCodes619
	ApproachSiteCodes620
	ApproachSiteCodes621
	ApproachSiteCodes622
	ApproachSiteCodes623
	ApproachSiteCodes624
	ApproachSiteCodes625
	ApproachSiteCodes626
	ApproachSiteCodes627
	ApproachSiteCodes628
	ApproachSiteCodes629
	ApproachSiteCodes630
	ApproachSiteCodes631
	ApproachSiteCodes632
	ApproachSiteCodes633
	ApproachSiteCodes634
	ApproachSiteCodes635
	ApproachSiteCodes636
	ApproachSiteCodes637
	ApproachSiteCodes638
	ApproachSiteCodes639
	ApproachSiteCodes640
	ApproachSiteCodes641
	ApproachSiteCodes642
	ApproachSiteCodes643
	ApproachSiteCodes644
	ApproachSiteCodes645
	ApproachSiteCodes646
	ApproachSiteCodes647
	ApproachSiteCodes648
	ApproachSiteCodes649
	ApproachSiteCodes650
	ApproachSiteCodes651
	ApproachSiteCodes652
	ApproachSiteCodes653
	ApproachSiteCodes654
	ApproachSiteCodes655
	ApproachSiteCodes656
	ApproachSiteCodes657
	ApproachSiteCodes658
	ApproachSiteCodes659
	ApproachSiteCodes660
	ApproachSiteCodes661
	ApproachSiteCodes662
	ApproachSiteCodes663
	ApproachSiteCodes664
	ApproachSiteCodes665
	ApproachSiteCodes666
	ApproachSiteCodes667
	ApproachSiteCodes668
	ApproachSiteCodes669
	ApproachSiteCodes670
	ApproachSiteCodes671
	ApproachSiteCodes672
	ApproachSiteCodes673
	ApproachSiteCodes674
	ApproachSiteCodes675
	ApproachSiteCodes676
	ApproachSiteCodes677
	ApproachSiteCodes678
	ApproachSiteCodes679
	ApproachSiteCodes680
	ApproachSiteCodes681
	ApproachSiteCodes682
	ApproachSiteCodes683
	ApproachSiteCodes684
	ApproachSiteCodes685
	ApproachSiteCodes686
	ApproachSiteCodes687
	ApproachSiteCodes688
	ApproachSiteCodes689
	ApproachSiteCodes690
	ApproachSiteCodes691
	ApproachSiteCodes692
	ApproachSiteCodes693
	ApproachSiteCodes694
	ApproachSiteCodes695
	ApproachSiteCodes696
	ApproachSiteCodes697
	ApproachSiteCodes698
	ApproachSiteCodes699
	ApproachSiteCodes700
	ApproachSiteCodes701
	ApproachSiteCodes702
	ApproachSiteCodes703
	ApproachSiteCodes704
	ApproachSiteCodes705
	ApproachSiteCodes706
	ApproachSiteCodes707
	ApproachSiteCodes708
	ApproachSiteCodes709
	ApproachSiteCodes710
	ApproachSiteCodes711
	ApproachSiteCodes712
	ApproachSiteCodes713
	ApproachSiteCodes714
	ApproachSiteCodes715
	ApproachSiteCodes716
	ApproachSiteCodes717
	ApproachSiteCodes718
	ApproachSiteCodes719
	ApproachSiteCodes720
	ApproachSiteCodes721
	ApproachSiteCodes722
	ApproachSiteCodes723
	ApproachSiteCodes724
	ApproachSiteCodes725
	ApproachSiteCodes726
	ApproachSiteCodes727
	ApproachSiteCodes728
	ApproachSiteCodes729
	ApproachSiteCodes730
	ApproachSiteCodes731
	ApproachSiteCodes732
	ApproachSiteCodes733
	ApproachSiteCodes734
	ApproachSiteCodes735
	ApproachSiteCodes736
	ApproachSiteCodes737
	ApproachSiteCodes738
	ApproachSiteCodes739
	ApproachSiteCodes740
	ApproachSiteCodes741
	ApproachSiteCodes742
	ApproachSiteCodes743
	ApproachSiteCodes744
	ApproachSiteCodes745
	ApproachSiteCodes746
	ApproachSiteCodes747
	ApproachSiteCodes748
	ApproachSiteCodes749
	ApproachSiteCodes750
	ApproachSiteCodes751
	ApproachSiteCodes752
	ApproachSiteCodes753
	ApproachSiteCodes754
	ApproachSiteCodes755
	ApproachSiteCodes756
	ApproachSiteCodes757
	ApproachSiteCodes758
	ApproachSiteCodes759
	ApproachSiteCodes760
	ApproachSiteCodes761
	ApproachSiteCodes762
	ApproachSiteCodes763
	ApproachSiteCodes764
	ApproachSiteCodes765
	ApproachSiteCodes766
	ApproachSiteCodes767
	ApproachSiteCodes768
	ApproachSiteCodes769
	ApproachSiteCodes770
	ApproachSiteCodes771
	ApproachSiteCodes772
	ApproachSiteCodes773
	ApproachSiteCodes774
	ApproachSiteCodes775
	ApproachSiteCodes776
	ApproachSiteCodes777
	ApproachSiteCodes778
	ApproachSiteCodes779
	ApproachSiteCodes780
	ApproachSiteCodes781
	ApproachSiteCodes782
	ApproachSiteCodes783
	ApproachSiteCodes784
	ApproachSiteCodes785
	ApproachSiteCodes786
	ApproachSiteCodes787
	ApproachSiteCodes788
	ApproachSiteCodes789
	ApproachSiteCodes790
	ApproachSiteCodes791
	ApproachSiteCodes792
	ApproachSiteCodes793
	ApproachSiteCodes794
	ApproachSiteCodes795
	ApproachSiteCodes796
	ApproachSiteCodes797
	ApproachSiteCodes798
	ApproachSiteCodes799
	ApproachSiteCodes800
	ApproachSiteCodes801
	ApproachSiteCodes802
	ApproachSiteCodes803
	ApproachSiteCodes804
	ApproachSiteCodes805
	ApproachSiteCodes806
	ApproachSiteCodes807
	ApproachSiteCodes808
	ApproachSiteCodes809
	ApproachSiteCodes810
	ApproachSiteCodes811
	ApproachSiteCodes812
	ApproachSiteCodes813
	ApproachSiteCodes814
	ApproachSiteCodes815
	ApproachSiteCodes816
	ApproachSiteCodes817
	ApproachSiteCodes818
	ApproachSiteCodes819
	ApproachSiteCodes820
	ApproachSiteCodes821
	ApproachSiteCodes822
	ApproachSiteCodes823
	ApproachSiteCodes824
	ApproachSiteCodes825
	ApproachSiteCodes826
	ApproachSiteCodes827
	ApproachSiteCodes828
	ApproachSiteCodes829
	ApproachSiteCodes830
	ApproachSiteCodes831
	ApproachSiteCodes832
	ApproachSiteCodes833
	ApproachSiteCodes834
	ApproachSiteCodes835
	ApproachSiteCodes836
	ApproachSiteCodes837
	ApproachSiteCodes838
	ApproachSiteCodes839
	ApproachSiteCodes840
	ApproachSiteCodes841
	ApproachSiteCodes842
	ApproachSiteCodes843
	ApproachSiteCodes844
	ApproachSiteCodes845
	ApproachSiteCodes846
	ApproachSiteCodes847
	ApproachSiteCodes848
	ApproachSiteCodes849
	ApproachSiteCodes850
	ApproachSiteCodes851
	ApproachSiteCodes852
	ApproachSiteCodes853
	ApproachSiteCodes854
	ApproachSiteCodes855
	ApproachSiteCodes856
	ApproachSiteCodes857
	ApproachSiteCodes858
	ApproachSiteCodes859
	ApproachSiteCodes860
	ApproachSiteCodes861
	ApproachSiteCodes862
	ApproachSiteCodes863
	ApproachSiteCodes864
	ApproachSiteCodes865
	ApproachSiteCodes866
	ApproachSiteCodes867
	ApproachSiteCodes868
	ApproachSiteCodes869
	ApproachSiteCodes870
	ApproachSiteCodes871
	ApproachSiteCodes872
	ApproachSiteCodes873
	ApproachSiteCodes874
	ApproachSiteCodes875
	ApproachSiteCodes876
	ApproachSiteCodes877
	ApproachSiteCodes878
	ApproachSiteCodes879
	ApproachSiteCodes880
	ApproachSiteCodes881
	ApproachSiteCodes882
	ApproachSiteCodes883
	ApproachSiteCodes884
	ApproachSiteCodes885
	ApproachSiteCodes886
	ApproachSiteCodes887
	ApproachSiteCodes888
	ApproachSiteCodes889
	ApproachSiteCodes890
	ApproachSiteCodes891
	ApproachSiteCodes892
	ApproachSiteCodes893
	ApproachSiteCodes894
	ApproachSiteCodes895
	ApproachSiteCodes896
	ApproachSiteCodes897
	ApproachSiteCodes898
	ApproachSiteCodes899
	ApproachSiteCodes900
	ApproachSiteCodes901
	ApproachSiteCodes902
	ApproachSiteCodes903
	ApproachSiteCodes904
	ApproachSiteCodes905
	ApproachSiteCodes906
	ApproachSiteCodes907
	ApproachSiteCodes908
	ApproachSiteCodes909
	ApproachSiteCodes910
	ApproachSiteCodes911
	ApproachSiteCodes912
	ApproachSiteCodes913
	ApproachSiteCodes914
	ApproachSiteCodes915
	ApproachSiteCodes916
	ApproachSiteCodes917
	ApproachSiteCodes918
	ApproachSiteCodes919
	ApproachSiteCodes920
	ApproachSiteCodes921
	ApproachSiteCodes922
	ApproachSiteCodes923
	ApproachSiteCodes924
	ApproachSiteCodes925
	ApproachSiteCodes926
	ApproachSiteCodes927
	ApproachSiteCodes928
	ApproachSiteCodes929
	ApproachSiteCodes930
	ApproachSiteCodes931
	ApproachSiteCodes932
	ApproachSiteCodes933
	ApproachSiteCodes934
	ApproachSiteCodes935
	ApproachSiteCodes936
	ApproachSiteCodes937
	ApproachSiteCodes938
	ApproachSiteCodes939
	ApproachSiteCodes940
	ApproachSiteCodes941
	ApproachSiteCodes942
	ApproachSiteCodes943
	ApproachSiteCodes944
	ApproachSiteCodes945
	ApproachSiteCodes946
	ApproachSiteCodes947
	ApproachSiteCodes948
	ApproachSiteCodes949
	ApproachSiteCodes950
	ApproachSiteCodes951
	ApproachSiteCodes952
	ApproachSiteCodes953
	ApproachSiteCodes954
	ApproachSiteCodes955
	ApproachSiteCodes956
	ApproachSiteCodes957
	ApproachSiteCodes958
	ApproachSiteCodes959
	ApproachSiteCodes960
	ApproachSiteCodes961
	ApproachSiteCodes962
	ApproachSiteCodes963
	ApproachSiteCodes964
	ApproachSiteCodes965
	ApproachSiteCodes966
	ApproachSiteCodes967
	ApproachSiteCodes968
	ApproachSiteCodes969
	ApproachSiteCodes970
	ApproachSiteCodes971
	ApproachSiteCodes972
	ApproachSiteCodes973
	ApproachSiteCodes974
	ApproachSiteCodes975
	ApproachSiteCodes976
	ApproachSiteCodes977
	ApproachSiteCodes978
	ApproachSiteCodes979
	ApproachSiteCodes980
	ApproachSiteCodes981
	ApproachSiteCodes982
	ApproachSiteCodes983
	ApproachSiteCodes984
	ApproachSiteCodes985
	ApproachSiteCodes986
	ApproachSiteCodes987
	ApproachSiteCodes988
	ApproachSiteCodes989
	ApproachSiteCodes990
	ApproachSiteCodes991
	ApproachSiteCodes992
	ApproachSiteCodes993
	ApproachSiteCodes994
	ApproachSiteCodes995
	ApproachSiteCodes996
	ApproachSiteCodes997
	ApproachSiteCodes998
	ApproachSiteCodes999
)

func AllApproachSiteCodes() []ApproachSiteCodes {
	return []ApproachSiteCodes{
		ApproachSiteCodes000,
		ApproachSiteCodes001,
		ApproachSiteCodes002,
		ApproachSiteCodes003,
		ApproachSiteCodes004,
		ApproachSiteCodes005,
		ApproachSiteCodes006,
		ApproachSiteCodes007,
		ApproachSiteCodes008,
		ApproachSiteCodes009,
		ApproachSiteCodes010,
		ApproachSiteCodes011,
		ApproachSiteCodes012,
		ApproachSiteCodes013,
		ApproachSiteCodes014,
		ApproachSiteCodes015,
		ApproachSiteCodes016,
		ApproachSiteCodes017,
		ApproachSiteCodes018,
		ApproachSiteCodes019,
		ApproachSiteCodes020,
		ApproachSiteCodes021,
		ApproachSiteCodes022,
		ApproachSiteCodes023,
		ApproachSiteCodes024,
		ApproachSiteCodes025,
		ApproachSiteCodes026,
		ApproachSiteCodes027,
		ApproachSiteCodes028,
		ApproachSiteCodes029,
		ApproachSiteCodes030,
		ApproachSiteCodes031,
		ApproachSiteCodes032,
		ApproachSiteCodes033,
		ApproachSiteCodes034,
		ApproachSiteCodes035,
		ApproachSiteCodes036,
		ApproachSiteCodes037,
		ApproachSiteCodes038,
		ApproachSiteCodes039,
		ApproachSiteCodes040,
		ApproachSiteCodes041,
		ApproachSiteCodes042,
		ApproachSiteCodes043,
		ApproachSiteCodes044,
		ApproachSiteCodes045,
		ApproachSiteCodes046,
		ApproachSiteCodes047,
		ApproachSiteCodes048,
		ApproachSiteCodes049,
		ApproachSiteCodes050,
		ApproachSiteCodes051,
		ApproachSiteCodes052,
		ApproachSiteCodes053,
		ApproachSiteCodes054,
		ApproachSiteCodes055,
		ApproachSiteCodes056,
		ApproachSiteCodes057,
		ApproachSiteCodes058,
		ApproachSiteCodes059,
		ApproachSiteCodes060,
		ApproachSiteCodes061,
		ApproachSiteCodes062,
		ApproachSiteCodes063,
		ApproachSiteCodes064,
		ApproachSiteCodes065,
		ApproachSiteCodes066,
		ApproachSiteCodes067,
		ApproachSiteCodes068,
		ApproachSiteCodes069,
		ApproachSiteCodes070,
		ApproachSiteCodes071,
		ApproachSiteCodes072,
		ApproachSiteCodes073,
		ApproachSiteCodes074,
		ApproachSiteCodes075,
		ApproachSiteCodes076,
		ApproachSiteCodes077,
		ApproachSiteCodes078,
		ApproachSiteCodes079,
		ApproachSiteCodes080,
		ApproachSiteCodes081,
		ApproachSiteCodes082,
		ApproachSiteCodes083,
		ApproachSiteCodes084,
		ApproachSiteCodes085,
		ApproachSiteCodes086,
		ApproachSiteCodes087,
		ApproachSiteCodes088,
		ApproachSiteCodes089,
		ApproachSiteCodes090,
		ApproachSiteCodes091,
		ApproachSiteCodes092,
		ApproachSiteCodes093,
		ApproachSiteCodes094,
		ApproachSiteCodes095,
		ApproachSiteCodes096,
		ApproachSiteCodes097,
		ApproachSiteCodes098,
		ApproachSiteCodes099,
		ApproachSiteCodes100,
		ApproachSiteCodes101,
		ApproachSiteCodes102,
		ApproachSiteCodes103,
		ApproachSiteCodes104,
		ApproachSiteCodes105,
		ApproachSiteCodes106,
		ApproachSiteCodes107,
		ApproachSiteCodes108,
		ApproachSiteCodes109,
		ApproachSiteCodes110,
		ApproachSiteCodes111,
		ApproachSiteCodes112,
		ApproachSiteCodes113,
		ApproachSiteCodes114,
		ApproachSiteCodes115,
		ApproachSiteCodes116,
		ApproachSiteCodes117,
		ApproachSiteCodes118,
		ApproachSiteCodes119,
		ApproachSiteCodes120,
		ApproachSiteCodes121,
		ApproachSiteCodes122,
		ApproachSiteCodes123,
		ApproachSiteCodes124,
		ApproachSiteCodes125,
		ApproachSiteCodes126,
		ApproachSiteCodes127,
		ApproachSiteCodes128,
		ApproachSiteCodes129,
		ApproachSiteCodes130,
		ApproachSiteCodes131,
		ApproachSiteCodes132,
		ApproachSiteCodes133,
		ApproachSiteCodes134,
		ApproachSiteCodes135,
		ApproachSiteCodes136,
		ApproachSiteCodes137,
		ApproachSiteCodes138,
		ApproachSiteCodes139,
		ApproachSiteCodes140,
		ApproachSiteCodes141,
		ApproachSiteCodes142,
		ApproachSiteCodes143,
		ApproachSiteCodes144,
		ApproachSiteCodes145,
		ApproachSiteCodes146,
		ApproachSiteCodes147,
		ApproachSiteCodes148,
		ApproachSiteCodes149,
		ApproachSiteCodes150,
		ApproachSiteCodes151,
		ApproachSiteCodes152,
		ApproachSiteCodes153,
		ApproachSiteCodes154,
		ApproachSiteCodes155,
		ApproachSiteCodes156,
		ApproachSiteCodes157,
		ApproachSiteCodes158,
		ApproachSiteCodes159,
		ApproachSiteCodes160,
		ApproachSiteCodes161,
		ApproachSiteCodes162,
		ApproachSiteCodes163,
		ApproachSiteCodes164,
		ApproachSiteCodes165,
		ApproachSiteCodes166,
		ApproachSiteCodes167,
		ApproachSiteCodes168,
		ApproachSiteCodes169,
		ApproachSiteCodes170,
		ApproachSiteCodes171,
		ApproachSiteCodes172,
		ApproachSiteCodes173,
		ApproachSiteCodes174,
		ApproachSiteCodes175,
		ApproachSiteCodes176,
		ApproachSiteCodes177,
		ApproachSiteCodes178,
		ApproachSiteCodes179,
		ApproachSiteCodes180,
		ApproachSiteCodes181,
		ApproachSiteCodes182,
		ApproachSiteCodes183,
		ApproachSiteCodes184,
		ApproachSiteCodes185,
		ApproachSiteCodes186,
		ApproachSiteCodes187,
		ApproachSiteCodes188,
		ApproachSiteCodes189,
		ApproachSiteCodes190,
		ApproachSiteCodes191,
		ApproachSiteCodes192,
		ApproachSiteCodes193,
		ApproachSiteCodes194,
		ApproachSiteCodes195,
		ApproachSiteCodes196,
		ApproachSiteCodes197,
		ApproachSiteCodes198,
		ApproachSiteCodes199,
		ApproachSiteCodes200,
		ApproachSiteCodes201,
		ApproachSiteCodes202,
		ApproachSiteCodes203,
		ApproachSiteCodes204,
		ApproachSiteCodes205,
		ApproachSiteCodes206,
		ApproachSiteCodes207,
		ApproachSiteCodes208,
		ApproachSiteCodes209,
		ApproachSiteCodes210,
		ApproachSiteCodes211,
		ApproachSiteCodes212,
		ApproachSiteCodes213,
		ApproachSiteCodes214,
		ApproachSiteCodes215,
		ApproachSiteCodes216,
		ApproachSiteCodes217,
		ApproachSiteCodes218,
		ApproachSiteCodes219,
		ApproachSiteCodes220,
		ApproachSiteCodes221,
		ApproachSiteCodes222,
		ApproachSiteCodes223,
		ApproachSiteCodes224,
		ApproachSiteCodes225,
		ApproachSiteCodes226,
		ApproachSiteCodes227,
		ApproachSiteCodes228,
		ApproachSiteCodes229,
		ApproachSiteCodes230,
		ApproachSiteCodes231,
		ApproachSiteCodes232,
		ApproachSiteCodes233,
		ApproachSiteCodes234,
		ApproachSiteCodes235,
		ApproachSiteCodes236,
		ApproachSiteCodes237,
		ApproachSiteCodes238,
		ApproachSiteCodes239,
		ApproachSiteCodes240,
		ApproachSiteCodes241,
		ApproachSiteCodes242,
		ApproachSiteCodes243,
		ApproachSiteCodes244,
		ApproachSiteCodes245,
		ApproachSiteCodes246,
		ApproachSiteCodes247,
		ApproachSiteCodes248,
		ApproachSiteCodes249,
		ApproachSiteCodes250,
		ApproachSiteCodes251,
		ApproachSiteCodes252,
		ApproachSiteCodes253,
		ApproachSiteCodes254,
		ApproachSiteCodes255,
		ApproachSiteCodes256,
		ApproachSiteCodes257,
		ApproachSiteCodes258,
		ApproachSiteCodes259,
		ApproachSiteCodes260,
		ApproachSiteCodes261,
		ApproachSiteCodes262,
		ApproachSiteCodes263,
		ApproachSiteCodes264,
		ApproachSiteCodes265,
		ApproachSiteCodes266,
		ApproachSiteCodes267,
		ApproachSiteCodes268,
		ApproachSiteCodes269,
		ApproachSiteCodes270,
		ApproachSiteCodes271,
		ApproachSiteCodes272,
		ApproachSiteCodes273,
		ApproachSiteCodes274,
		ApproachSiteCodes275,
		ApproachSiteCodes276,
		ApproachSiteCodes277,
		ApproachSiteCodes278,
		ApproachSiteCodes279,
		ApproachSiteCodes280,
		ApproachSiteCodes281,
		ApproachSiteCodes282,
		ApproachSiteCodes283,
		ApproachSiteCodes284,
		ApproachSiteCodes285,
		ApproachSiteCodes286,
		ApproachSiteCodes287,
		ApproachSiteCodes288,
		ApproachSiteCodes289,
		ApproachSiteCodes290,
		ApproachSiteCodes291,
		ApproachSiteCodes292,
		ApproachSiteCodes293,
		ApproachSiteCodes294,
		ApproachSiteCodes295,
		ApproachSiteCodes296,
		ApproachSiteCodes297,
		ApproachSiteCodes298,
		ApproachSiteCodes299,
		ApproachSiteCodes300,
		ApproachSiteCodes301,
		ApproachSiteCodes302,
		ApproachSiteCodes303,
		ApproachSiteCodes304,
		ApproachSiteCodes305,
		ApproachSiteCodes306,
		ApproachSiteCodes307,
		ApproachSiteCodes308,
		ApproachSiteCodes309,
		ApproachSiteCodes310,
		ApproachSiteCodes311,
		ApproachSiteCodes312,
		ApproachSiteCodes313,
		ApproachSiteCodes314,
		ApproachSiteCodes315,
		ApproachSiteCodes316,
		ApproachSiteCodes317,
		ApproachSiteCodes318,
		ApproachSiteCodes319,
		ApproachSiteCodes320,
		ApproachSiteCodes321,
		ApproachSiteCodes322,
		ApproachSiteCodes323,
		ApproachSiteCodes324,
		ApproachSiteCodes325,
		ApproachSiteCodes326,
		ApproachSiteCodes327,
		ApproachSiteCodes328,
		ApproachSiteCodes329,
		ApproachSiteCodes330,
		ApproachSiteCodes331,
		ApproachSiteCodes332,
		ApproachSiteCodes333,
		ApproachSiteCodes334,
		ApproachSiteCodes335,
		ApproachSiteCodes336,
		ApproachSiteCodes337,
		ApproachSiteCodes338,
		ApproachSiteCodes339,
		ApproachSiteCodes340,
		ApproachSiteCodes341,
		ApproachSiteCodes342,
		ApproachSiteCodes343,
		ApproachSiteCodes344,
		ApproachSiteCodes345,
		ApproachSiteCodes346,
		ApproachSiteCodes347,
		ApproachSiteCodes348,
		ApproachSiteCodes349,
		ApproachSiteCodes350,
		ApproachSiteCodes351,
		ApproachSiteCodes352,
		ApproachSiteCodes353,
		ApproachSiteCodes354,
		ApproachSiteCodes355,
		ApproachSiteCodes356,
		ApproachSiteCodes357,
		ApproachSiteCodes358,
		ApproachSiteCodes359,
		ApproachSiteCodes360,
		ApproachSiteCodes361,
		ApproachSiteCodes362,
		ApproachSiteCodes363,
		ApproachSiteCodes364,
		ApproachSiteCodes365,
		ApproachSiteCodes366,
		ApproachSiteCodes367,
		ApproachSiteCodes368,
		ApproachSiteCodes369,
		ApproachSiteCodes370,
		ApproachSiteCodes371,
		ApproachSiteCodes372,
		ApproachSiteCodes373,
		ApproachSiteCodes374,
		ApproachSiteCodes375,
		ApproachSiteCodes376,
		ApproachSiteCodes377,
		ApproachSiteCodes378,
		ApproachSiteCodes379,
		ApproachSiteCodes380,
		ApproachSiteCodes381,
		ApproachSiteCodes382,
		ApproachSiteCodes383,
		ApproachSiteCodes384,
		ApproachSiteCodes385,
		ApproachSiteCodes386,
		ApproachSiteCodes387,
		ApproachSiteCodes388,
		ApproachSiteCodes389,
		ApproachSiteCodes390,
		ApproachSiteCodes391,
		ApproachSiteCodes392,
		ApproachSiteCodes393,
		ApproachSiteCodes394,
		ApproachSiteCodes395,
		ApproachSiteCodes396,
		ApproachSiteCodes397,
		ApproachSiteCodes398,
		ApproachSiteCodes399,
		ApproachSiteCodes400,
		ApproachSiteCodes401,
		ApproachSiteCodes402,
		ApproachSiteCodes403,
		ApproachSiteCodes404,
		ApproachSiteCodes405,
		ApproachSiteCodes406,
		ApproachSiteCodes407,
		ApproachSiteCodes408,
		ApproachSiteCodes409,
		ApproachSiteCodes410,
		ApproachSiteCodes411,
		ApproachSiteCodes412,
		ApproachSiteCodes413,
		ApproachSiteCodes414,
		ApproachSiteCodes415,
		ApproachSiteCodes416,
		ApproachSiteCodes417,
		ApproachSiteCodes418,
		ApproachSiteCodes419,
		ApproachSiteCodes420,
		ApproachSiteCodes421,
		ApproachSiteCodes422,
		ApproachSiteCodes423,
		ApproachSiteCodes424,
		ApproachSiteCodes425,
		ApproachSiteCodes426,
		ApproachSiteCodes427,
		ApproachSiteCodes428,
		ApproachSiteCodes429,
		ApproachSiteCodes430,
		ApproachSiteCodes431,
		ApproachSiteCodes432,
		ApproachSiteCodes433,
		ApproachSiteCodes434,
		ApproachSiteCodes435,
		ApproachSiteCodes436,
		ApproachSiteCodes437,
		ApproachSiteCodes438,
		ApproachSiteCodes439,
		ApproachSiteCodes440,
		ApproachSiteCodes441,
		ApproachSiteCodes442,
		ApproachSiteCodes443,
		ApproachSiteCodes444,
		ApproachSiteCodes445,
		ApproachSiteCodes446,
		ApproachSiteCodes447,
		ApproachSiteCodes448,
		ApproachSiteCodes449,
		ApproachSiteCodes450,
		ApproachSiteCodes451,
		ApproachSiteCodes452,
		ApproachSiteCodes453,
		ApproachSiteCodes454,
		ApproachSiteCodes455,
		ApproachSiteCodes456,
		ApproachSiteCodes457,
		ApproachSiteCodes458,
		ApproachSiteCodes459,
		ApproachSiteCodes460,
		ApproachSiteCodes461,
		ApproachSiteCodes462,
		ApproachSiteCodes463,
		ApproachSiteCodes464,
		ApproachSiteCodes465,
		ApproachSiteCodes466,
		ApproachSiteCodes467,
		ApproachSiteCodes468,
		ApproachSiteCodes469,
		ApproachSiteCodes470,
		ApproachSiteCodes471,
		ApproachSiteCodes472,
		ApproachSiteCodes473,
		ApproachSiteCodes474,
		ApproachSiteCodes475,
		ApproachSiteCodes476,
		ApproachSiteCodes477,
		ApproachSiteCodes478,
		ApproachSiteCodes479,
		ApproachSiteCodes480,
		ApproachSiteCodes481,
		ApproachSiteCodes482,
		ApproachSiteCodes483,
		ApproachSiteCodes484,
		ApproachSiteCodes485,
		ApproachSiteCodes486,
		ApproachSiteCodes487,
		ApproachSiteCodes488,
		ApproachSiteCodes489,
		ApproachSiteCodes490,
		ApproachSiteCodes491,
		ApproachSiteCodes492,
		ApproachSiteCodes493,
		ApproachSiteCodes494,
		ApproachSiteCodes495,
		ApproachSiteCodes496,
		ApproachSiteCodes497,
		ApproachSiteCodes498,
		ApproachSiteCodes499,
		ApproachSiteCodes500,
		ApproachSiteCodes501,
		ApproachSiteCodes502,
		ApproachSiteCodes503,
		ApproachSiteCodes504,
		ApproachSiteCodes505,
		ApproachSiteCodes506,
		ApproachSiteCodes507,
		ApproachSiteCodes508,
		ApproachSiteCodes509,
		ApproachSiteCodes510,
		ApproachSiteCodes511,
		ApproachSiteCodes512,
		ApproachSiteCodes513,
		ApproachSiteCodes514,
		ApproachSiteCodes515,
		ApproachSiteCodes516,
		ApproachSiteCodes517,
		ApproachSiteCodes518,
		ApproachSiteCodes519,
		ApproachSiteCodes520,
		ApproachSiteCodes521,
		ApproachSiteCodes522,
		ApproachSiteCodes523,
		ApproachSiteCodes524,
		ApproachSiteCodes525,
		ApproachSiteCodes526,
		ApproachSiteCodes527,
		ApproachSiteCodes528,
		ApproachSiteCodes529,
		ApproachSiteCodes530,
		ApproachSiteCodes531,
		ApproachSiteCodes532,
		ApproachSiteCodes533,
		ApproachSiteCodes534,
		ApproachSiteCodes535,
		ApproachSiteCodes536,
		ApproachSiteCodes537,
		ApproachSiteCodes538,
		ApproachSiteCodes539,
		ApproachSiteCodes540,
		ApproachSiteCodes541,
		ApproachSiteCodes542,
		ApproachSiteCodes543,
		ApproachSiteCodes544,
		ApproachSiteCodes545,
		ApproachSiteCodes546,
		ApproachSiteCodes547,
		ApproachSiteCodes548,
		ApproachSiteCodes549,
		ApproachSiteCodes550,
		ApproachSiteCodes551,
		ApproachSiteCodes552,
		ApproachSiteCodes553,
		ApproachSiteCodes554,
		ApproachSiteCodes555,
		ApproachSiteCodes556,
		ApproachSiteCodes557,
		ApproachSiteCodes558,
		ApproachSiteCodes559,
		ApproachSiteCodes560,
		ApproachSiteCodes561,
		ApproachSiteCodes562,
		ApproachSiteCodes563,
		ApproachSiteCodes564,
		ApproachSiteCodes565,
		ApproachSiteCodes566,
		ApproachSiteCodes567,
		ApproachSiteCodes568,
		ApproachSiteCodes569,
		ApproachSiteCodes570,
		ApproachSiteCodes571,
		ApproachSiteCodes572,
		ApproachSiteCodes573,
		ApproachSiteCodes574,
		ApproachSiteCodes575,
		ApproachSiteCodes576,
		ApproachSiteCodes577,
		ApproachSiteCodes578,
		ApproachSiteCodes579,
		ApproachSiteCodes580,
		ApproachSiteCodes581,
		ApproachSiteCodes582,
		ApproachSiteCodes583,
		ApproachSiteCodes584,
		ApproachSiteCodes585,
		ApproachSiteCodes586,
		ApproachSiteCodes587,
		ApproachSiteCodes588,
		ApproachSiteCodes589,
		ApproachSiteCodes590,
		ApproachSiteCodes591,
		ApproachSiteCodes592,
		ApproachSiteCodes593,
		ApproachSiteCodes594,
		ApproachSiteCodes595,
		ApproachSiteCodes596,
		ApproachSiteCodes597,
		ApproachSiteCodes598,
		ApproachSiteCodes599,
		ApproachSiteCodes600,
		ApproachSiteCodes601,
		ApproachSiteCodes602,
		ApproachSiteCodes603,
		ApproachSiteCodes604,
		ApproachSiteCodes605,
		ApproachSiteCodes606,
		ApproachSiteCodes607,
		ApproachSiteCodes608,
		ApproachSiteCodes609,
		ApproachSiteCodes610,
		ApproachSiteCodes611,
		ApproachSiteCodes612,
		ApproachSiteCodes613,
		ApproachSiteCodes614,
		ApproachSiteCodes615,
		ApproachSiteCodes616,
		ApproachSiteCodes617,
		ApproachSiteCodes618,
		ApproachSiteCodes619,
		ApproachSiteCodes620,
		ApproachSiteCodes621,
		ApproachSiteCodes622,
		ApproachSiteCodes623,
		ApproachSiteCodes624,
		ApproachSiteCodes625,
		ApproachSiteCodes626,
		ApproachSiteCodes627,
		ApproachSiteCodes628,
		ApproachSiteCodes629,
		ApproachSiteCodes630,
		ApproachSiteCodes631,
		ApproachSiteCodes632,
		ApproachSiteCodes633,
		ApproachSiteCodes634,
		ApproachSiteCodes635,
		ApproachSiteCodes636,
		ApproachSiteCodes637,
		ApproachSiteCodes638,
		ApproachSiteCodes639,
		ApproachSiteCodes640,
		ApproachSiteCodes641,
		ApproachSiteCodes642,
		ApproachSiteCodes643,
		ApproachSiteCodes644,
		ApproachSiteCodes645,
		ApproachSiteCodes646,
		ApproachSiteCodes647,
		ApproachSiteCodes648,
		ApproachSiteCodes649,
		ApproachSiteCodes650,
		ApproachSiteCodes651,
		ApproachSiteCodes652,
		ApproachSiteCodes653,
		ApproachSiteCodes654,
		ApproachSiteCodes655,
		ApproachSiteCodes656,
		ApproachSiteCodes657,
		ApproachSiteCodes658,
		ApproachSiteCodes659,
		ApproachSiteCodes660,
		ApproachSiteCodes661,
		ApproachSiteCodes662,
		ApproachSiteCodes663,
		ApproachSiteCodes664,
		ApproachSiteCodes665,
		ApproachSiteCodes666,
		ApproachSiteCodes667,
		ApproachSiteCodes668,
		ApproachSiteCodes669,
		ApproachSiteCodes670,
		ApproachSiteCodes671,
		ApproachSiteCodes672,
		ApproachSiteCodes673,
		ApproachSiteCodes674,
		ApproachSiteCodes675,
		ApproachSiteCodes676,
		ApproachSiteCodes677,
		ApproachSiteCodes678,
		ApproachSiteCodes679,
		ApproachSiteCodes680,
		ApproachSiteCodes681,
		ApproachSiteCodes682,
		ApproachSiteCodes683,
		ApproachSiteCodes684,
		ApproachSiteCodes685,
		ApproachSiteCodes686,
		ApproachSiteCodes687,
		ApproachSiteCodes688,
		ApproachSiteCodes689,
		ApproachSiteCodes690,
		ApproachSiteCodes691,
		ApproachSiteCodes692,
		ApproachSiteCodes693,
		ApproachSiteCodes694,
		ApproachSiteCodes695,
		ApproachSiteCodes696,
		ApproachSiteCodes697,
		ApproachSiteCodes698,
		ApproachSiteCodes699,
		ApproachSiteCodes700,
		ApproachSiteCodes701,
		ApproachSiteCodes702,
		ApproachSiteCodes703,
		ApproachSiteCodes704,
		ApproachSiteCodes705,
		ApproachSiteCodes706,
		ApproachSiteCodes707,
		ApproachSiteCodes708,
		ApproachSiteCodes709,
		ApproachSiteCodes710,
		ApproachSiteCodes711,
		ApproachSiteCodes712,
		ApproachSiteCodes713,
		ApproachSiteCodes714,
		ApproachSiteCodes715,
		ApproachSiteCodes716,
		ApproachSiteCodes717,
		ApproachSiteCodes718,
		ApproachSiteCodes719,
		ApproachSiteCodes720,
		ApproachSiteCodes721,
		ApproachSiteCodes722,
		ApproachSiteCodes723,
		ApproachSiteCodes724,
		ApproachSiteCodes725,
		ApproachSiteCodes726,
		ApproachSiteCodes727,
		ApproachSiteCodes728,
		ApproachSiteCodes729,
		ApproachSiteCodes730,
		ApproachSiteCodes731,
		ApproachSiteCodes732,
		ApproachSiteCodes733,
		ApproachSiteCodes734,
		ApproachSiteCodes735,
		ApproachSiteCodes736,
		ApproachSiteCodes737,
		ApproachSiteCodes738,
		ApproachSiteCodes739,
		ApproachSiteCodes740,
		ApproachSiteCodes741,
		ApproachSiteCodes742,
		ApproachSiteCodes743,
		ApproachSiteCodes744,
		ApproachSiteCodes745,
		ApproachSiteCodes746,
		ApproachSiteCodes747,
		ApproachSiteCodes748,
		ApproachSiteCodes749,
		ApproachSiteCodes750,
		ApproachSiteCodes751,
		ApproachSiteCodes752,
		ApproachSiteCodes753,
		ApproachSiteCodes754,
		ApproachSiteCodes755,
		ApproachSiteCodes756,
		ApproachSiteCodes757,
		ApproachSiteCodes758,
		ApproachSiteCodes759,
		ApproachSiteCodes760,
		ApproachSiteCodes761,
		ApproachSiteCodes762,
		ApproachSiteCodes763,
		ApproachSiteCodes764,
		ApproachSiteCodes765,
		ApproachSiteCodes766,
		ApproachSiteCodes767,
		ApproachSiteCodes768,
		ApproachSiteCodes769,
		ApproachSiteCodes770,
		ApproachSiteCodes771,
		ApproachSiteCodes772,
		ApproachSiteCodes773,
		ApproachSiteCodes774,
		ApproachSiteCodes775,
		ApproachSiteCodes776,
		ApproachSiteCodes777,
		ApproachSiteCodes778,
		ApproachSiteCodes779,
		ApproachSiteCodes780,
		ApproachSiteCodes781,
		ApproachSiteCodes782,
		ApproachSiteCodes783,
		ApproachSiteCodes784,
		ApproachSiteCodes785,
		ApproachSiteCodes786,
		ApproachSiteCodes787,
		ApproachSiteCodes788,
		ApproachSiteCodes789,
		ApproachSiteCodes790,
		ApproachSiteCodes791,
		ApproachSiteCodes792,
		ApproachSiteCodes793,
		ApproachSiteCodes794,
		ApproachSiteCodes795,
		ApproachSiteCodes796,
		ApproachSiteCodes797,
		ApproachSiteCodes798,
		ApproachSiteCodes799,
		ApproachSiteCodes800,
		ApproachSiteCodes801,
		ApproachSiteCodes802,
		ApproachSiteCodes803,
		ApproachSiteCodes804,
		ApproachSiteCodes805,
		ApproachSiteCodes806,
		ApproachSiteCodes807,
		ApproachSiteCodes808,
		ApproachSiteCodes809,
		ApproachSiteCodes810,
		ApproachSiteCodes811,
		ApproachSiteCodes812,
		ApproachSiteCodes813,
		ApproachSiteCodes814,
		ApproachSiteCodes815,
		ApproachSiteCodes816,
		ApproachSiteCodes817,
		ApproachSiteCodes818,
		ApproachSiteCodes819,
		ApproachSiteCodes820,
		ApproachSiteCodes821,
		ApproachSiteCodes822,
		ApproachSiteCodes823,
		ApproachSiteCodes824,
		ApproachSiteCodes825,
		ApproachSiteCodes826,
		ApproachSiteCodes827,
		ApproachSiteCodes828,
		ApproachSiteCodes829,
		ApproachSiteCodes830,
		ApproachSiteCodes831,
		ApproachSiteCodes832,
		ApproachSiteCodes833,
		ApproachSiteCodes834,
		ApproachSiteCodes835,
		ApproachSiteCodes836,
		ApproachSiteCodes837,
		ApproachSiteCodes838,
		ApproachSiteCodes839,
		ApproachSiteCodes840,
		ApproachSiteCodes841,
		ApproachSiteCodes842,
		ApproachSiteCodes843,
		ApproachSiteCodes844,
		ApproachSiteCodes845,
		ApproachSiteCodes846,
		ApproachSiteCodes847,
		ApproachSiteCodes848,
		ApproachSiteCodes849,
		ApproachSiteCodes850,
		ApproachSiteCodes851,
		ApproachSiteCodes852,
		ApproachSiteCodes853,
		ApproachSiteCodes854,
		ApproachSiteCodes855,
		ApproachSiteCodes856,
		ApproachSiteCodes857,
		ApproachSiteCodes858,
		ApproachSiteCodes859,
		ApproachSiteCodes860,
		ApproachSiteCodes861,
		ApproachSiteCodes862,
		ApproachSiteCodes863,
		ApproachSiteCodes864,
		ApproachSiteCodes865,
		ApproachSiteCodes866,
		ApproachSiteCodes867,
		ApproachSiteCodes868,
		ApproachSiteCodes869,
		ApproachSiteCodes870,
		ApproachSiteCodes871,
		ApproachSiteCodes872,
		ApproachSiteCodes873,
		ApproachSiteCodes874,
		ApproachSiteCodes875,
		ApproachSiteCodes876,
		ApproachSiteCodes877,
		ApproachSiteCodes878,
		ApproachSiteCodes879,
		ApproachSiteCodes880,
		ApproachSiteCodes881,
		ApproachSiteCodes882,
		ApproachSiteCodes883,
		ApproachSiteCodes884,
		ApproachSiteCodes885,
		ApproachSiteCodes886,
		ApproachSiteCodes887,
		ApproachSiteCodes888,
		ApproachSiteCodes889,
		ApproachSiteCodes890,
		ApproachSiteCodes891,
		ApproachSiteCodes892,
		ApproachSiteCodes893,
		ApproachSiteCodes894,
		ApproachSiteCodes895,
		ApproachSiteCodes896,
		ApproachSiteCodes897,
		ApproachSiteCodes898,
		ApproachSiteCodes899,
		ApproachSiteCodes900,
		ApproachSiteCodes901,
		ApproachSiteCodes902,
		ApproachSiteCodes903,
		ApproachSiteCodes904,
		ApproachSiteCodes905,
		ApproachSiteCodes906,
		ApproachSiteCodes907,
		ApproachSiteCodes908,
		ApproachSiteCodes909,
		ApproachSiteCodes910,
		ApproachSiteCodes911,
		ApproachSiteCodes912,
		ApproachSiteCodes913,
		ApproachSiteCodes914,
		ApproachSiteCodes915,
		ApproachSiteCodes916,
		ApproachSiteCodes917,
		ApproachSiteCodes918,
		ApproachSiteCodes919,
		ApproachSiteCodes920,
		ApproachSiteCodes921,
		ApproachSiteCodes922,
		ApproachSiteCodes923,
		ApproachSiteCodes924,
		ApproachSiteCodes925,
		ApproachSiteCodes926,
		ApproachSiteCodes927,
		ApproachSiteCodes928,
		ApproachSiteCodes929,
		ApproachSiteCodes930,
		ApproachSiteCodes931,
		ApproachSiteCodes932,
		ApproachSiteCodes933,
		ApproachSiteCodes934,
		ApproachSiteCodes935,
		ApproachSiteCodes936,
		ApproachSiteCodes937,
		ApproachSiteCodes938,
		ApproachSiteCodes939,
		ApproachSiteCodes940,
		ApproachSiteCodes941,
		ApproachSiteCodes942,
		ApproachSiteCodes943,
		ApproachSiteCodes944,
		ApproachSiteCodes945,
		ApproachSiteCodes946,
		ApproachSiteCodes947,
		ApproachSiteCodes948,
		ApproachSiteCodes949,
		ApproachSiteCodes950,
		ApproachSiteCodes951,
		ApproachSiteCodes952,
		ApproachSiteCodes953,
		ApproachSiteCodes954,
		ApproachSiteCodes955,
		ApproachSiteCodes956,
		ApproachSiteCodes957,
		ApproachSiteCodes958,
		ApproachSiteCodes959,
		ApproachSiteCodes960,
		ApproachSiteCodes961,
		ApproachSiteCodes962,
		ApproachSiteCodes963,
		ApproachSiteCodes964,
		ApproachSiteCodes965,
		ApproachSiteCodes966,
		ApproachSiteCodes967,
		ApproachSiteCodes968,
		ApproachSiteCodes969,
		ApproachSiteCodes970,
		ApproachSiteCodes971,
		ApproachSiteCodes972,
		ApproachSiteCodes973,
		ApproachSiteCodes974,
		ApproachSiteCodes975,
		ApproachSiteCodes976,
		ApproachSiteCodes977,
		ApproachSiteCodes978,
		ApproachSiteCodes979,
		ApproachSiteCodes980,
		ApproachSiteCodes981,
		ApproachSiteCodes982,
		ApproachSiteCodes983,
		ApproachSiteCodes984,
		ApproachSiteCodes985,
		ApproachSiteCodes986,
		ApproachSiteCodes987,
		ApproachSiteCodes988,
		ApproachSiteCodes989,
		ApproachSiteCodes990,
		ApproachSiteCodes991,
		ApproachSiteCodes992,
		ApproachSiteCodes993,
		ApproachSiteCodes994,
		ApproachSiteCodes995,
		ApproachSiteCodes996,
		ApproachSiteCodes997,
		ApproachSiteCodes998,
		ApproachSiteCodes999,
	}
}

func FindApproachSiteCodes(filter string) []ApproachSiteCodes {
	ret := make([]ApproachSiteCodes, 0)
	for _, at := range AllApproachSiteCodes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ApproachSiteCodes) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ApproachSiteCodes) String() string {
	switch cpt {
	case ApproachSiteCodes000:
		return "Anatomical structure"
	case ApproachSiteCodes001:
		return "Posterior carpal region"
	case ApproachSiteCodes002:
		return "Fetal part of placenta"
	case ApproachSiteCodes003:
		return "Entire condylar emissary vein"
	case ApproachSiteCodes004:
		return "Visceral layer of Bowman's capsule"
	case ApproachSiteCodes005:
		return "Parathyroid gland"
	case ApproachSiteCodes006:
		return "Subcutaneous tissue of medial surface of index finger"
	case ApproachSiteCodes007:
		return "Coronoid process of mandible"
	case ApproachSiteCodes008:
		return "Central pair of microtubules, cilium or flagellum, not bacterial"
	case ApproachSiteCodes009:
		return "Deep circumflex artery of ilium"
	case ApproachSiteCodes010:
		return "Supraclavicular part of brachial plexus"
	case ApproachSiteCodes011:
		return "Anterior division of renal artery"
	case ApproachSiteCodes012:
		return "Entire left commissure of aortic valve"
	case ApproachSiteCodes013:
		return "Gluteus maximus muscle"
	case ApproachSiteCodes014:
		return "Articular surface, phalanges, of fourth metacarpal bone"
	case ApproachSiteCodes015:
		return "Canal of Hering"
	case ApproachSiteCodes016:
		return "Hepatocolic ligament"
	case ApproachSiteCodes017:
		return "Superior labial artery"
	case ApproachSiteCodes018:
		return "Lateral vestibular nucleus"
	case ApproachSiteCodes019:
		return "Mesotympanum"
	case ApproachSiteCodes020:
		return "Pectoral region"
	case ApproachSiteCodes021:
		return "Kupffer cell"
	case ApproachSiteCodes022:
		return "Thoracic nerve"
	case ApproachSiteCodes023:
		return "Right lower lobe of lung"
	case ApproachSiteCodes024:
		return "Superior articular process of lumbar vertebra"
	case ApproachSiteCodes025:
		return "Lateral myocardium"
	case ApproachSiteCodes026:
		return "Central axillary lymph node"
	case ApproachSiteCodes027:
		return "Entire flexor tendon and tendon sheath of fourth toe"
	case ApproachSiteCodes028:
		return "Metacarpophalangeal joint of index finger"
	case ApproachSiteCodes029:
		return "Fifth metatarsal bone"
	case ApproachSiteCodes030:
		return "Plantar surface of great toe"
	case ApproachSiteCodes031:
		return "Skin of umbilicus"
	case ApproachSiteCodes032:
		return "Cardiac impression of liver"
	case ApproachSiteCodes033:
		return "Ankle"
	case ApproachSiteCodes034:
		return "Atrioventricular bundle"
	case ApproachSiteCodes035:
		return "Reticular corium"
	case ApproachSiteCodes036:
		return "Wall of urinary bladder"
	case ApproachSiteCodes037:
		return "Dental branches of inferior alveolar artery"
	case ApproachSiteCodes038:
		return "Entire posterior temporal diploic vein"
	case ApproachSiteCodes039:
		return "Gastric fundus"
	case ApproachSiteCodes040:
		return "Inferior surface of tongue"
	case ApproachSiteCodes041:
		return "Trochanteric bursa"
	case ApproachSiteCodes042:
		return "Collateral ligament"
	case ApproachSiteCodes043:
		return "Lateral corticospinal tract"
	case ApproachSiteCodes044:
		return "Basophilic normoblast"
	case ApproachSiteCodes045:
		return "Ascending frontal gyrus"
	case ApproachSiteCodes046:
		return "Flexor hallucis longus muscle"
	case ApproachSiteCodes047:
		return "Cardiopulmonary circulatory system"
	case ApproachSiteCodes048:
		return "Transverse colon"
	case ApproachSiteCodes049:
		return "Costal surface of lung"
	case ApproachSiteCodes050:
		return "Entire vagus nerve parasympathetic fibers to cardiac plexus"
	case ApproachSiteCodes051:
		return "Intervertebral disc space of fifth lumbar vertebra"
	case ApproachSiteCodes052:
		return "Head of phalanx of great toe"
	case ApproachSiteCodes053:
		return "Capsule of proximal interphalangeal joint of third toe"
	case ApproachSiteCodes054:
		return "Interventricular septum"
	case ApproachSiteCodes055:
		return "Palpebral fissure"
	case ApproachSiteCodes056:
		return "Subcutaneous tissue of philtrum"
	case ApproachSiteCodes057:
		return "Multivesicular body, internal vesicles"
	case ApproachSiteCodes058:
		return "Tuberosity of distal phalanx of little toe"
	case ApproachSiteCodes059:
		return "Superior articular process of seventh thoracic vertebra"
	case ApproachSiteCodes060:
		return "Tracheal mucous membrane"
	case ApproachSiteCodes061:
		return "Jaw"
	case ApproachSiteCodes062:
		return "Embryonic structure"
	case ApproachSiteCodes063:
		return "Fetal hyaloid artery"
	case ApproachSiteCodes064:
		return "Small intestine submucosa"
	case ApproachSiteCodes065:
		return "Body of ischium"
	case ApproachSiteCodes066:
		return "Dense intermediate filament bundles"
	case ApproachSiteCodes067:
		return "Head and neck structure"
	case ApproachSiteCodes068:
		return "Visceral aspect of liver"
	case ApproachSiteCodes069:
		return "Entire deep temporal vein"
	case ApproachSiteCodes070:
		return "Posterior intercostal arteries"
	case ApproachSiteCodes071:
		return "Fetal chondrocranium"
	case ApproachSiteCodes072:
		return "Posterior cervical spinal cord nerve root"
	case ApproachSiteCodes073:
		return "Spinous process of fifth thoracic vertebra"
	case ApproachSiteCodes074:
		return "Oral region of face"
	case ApproachSiteCodes075:
		return "Lamina muscularis of colonic mucous membrane"
	case ApproachSiteCodes076:
		return "Anterior cruciate ligament of knee joint"
	case ApproachSiteCodes077:
		return "Superior laryngeal aperture"
	case ApproachSiteCodes078:
		return "Thyrohyoid branch of ansa cervicalis"
	case ApproachSiteCodes079:
		return "Crus of diaphragm"
	case ApproachSiteCodes080:
		return "Bronchus"
	case ApproachSiteCodes081:
		return "Ovarian vein"
	case ApproachSiteCodes082:
		return "Meningeal branches of occipital artery"
	case ApproachSiteCodes083:
		return "Entire diaphragmatic lymph node"
	case ApproachSiteCodes084:
		return "Fibrous portion of pericardium"
	case ApproachSiteCodes085:
		return "Peritonsillar tissue"
	case ApproachSiteCodes086:
		return "Sebaceous gland"
	case ApproachSiteCodes087:
		return "Vesicular bursa of sternohyoid muscle"
	case ApproachSiteCodes088:
		return "Frontozygomatic suture of skull"
	case ApproachSiteCodes089:
		return "Promonocyte"
	case ApproachSiteCodes090:
		return "Entire subcutaneous prepatellar bursa"
	case ApproachSiteCodes091:
		return "Female"
	case ApproachSiteCodes092:
		return "Sternothyroid muscle"
	case ApproachSiteCodes093:
		return "Superior occipital gyrus"
	case ApproachSiteCodes094:
		return "Thymic cortex"
	case ApproachSiteCodes095:
		return "Cranial cavity"
	case ApproachSiteCodes096:
		return "Entire major calyx"
	case ApproachSiteCodes097:
		return "Tarsal gland"
	case ApproachSiteCodes098:
		return "Inferior longitudinal muscle of tongue"
	case ApproachSiteCodes099:
		return "Aortopulmonary septum"
	case ApproachSiteCodes100:
		return "Frenulum linguae"
	case ApproachSiteCodes101:
		return "Odontoid process of axis"
	case ApproachSiteCodes102:
		return "Mandibular nerve"
	case ApproachSiteCodes103:
		return "Chromosomes, group E"
	case ApproachSiteCodes104:
		return "Teres major muscle"
	case ApproachSiteCodes105:
		return "Synostosis"
	case ApproachSiteCodes106:
		return "Meninges"
	case ApproachSiteCodes107:
		return "Duodenal serosa"
	case ApproachSiteCodes108:
		return "Inferior articular process of sixth cervical vertebra"
	case ApproachSiteCodes109:
		return "Dorsal digital nerves of radial nerve"
	case ApproachSiteCodes110:
		return "Distinctive arrangement of microtubules"
	case ApproachSiteCodes111:
		return "Vertebral nerve"
	case ApproachSiteCodes112:
		return "Glottis"
	case ApproachSiteCodes113:
		return "Telogen hair"
	case ApproachSiteCodes114:
		return "Deep flexor tendon of index finger"
	case ApproachSiteCodes115:
		return "Gastric serosa"
	case ApproachSiteCodes116:
		return "Vastus lateralis muscle"
	case ApproachSiteCodes117:
		return "Posterior limb of stapes"
	case ApproachSiteCodes118:
		return "Paravesicular lymph node"
	case ApproachSiteCodes119:
		return "Ventricle of larynx"
	case ApproachSiteCodes120:
		return "Yellow fibrocartilage"
	case ApproachSiteCodes121:
		return "Parietal branch of superficial temporal artery"
	case ApproachSiteCodes122:
		return "Metatarsus"
	case ApproachSiteCodes123:
		return "Soft tissues of trunk"
	case ApproachSiteCodes124:
		return "Anterior cecal artery"
	case ApproachSiteCodes125:
		return "Ejaculatory duct"
	case ApproachSiteCodes126:
		return "Frontomental diameter of head"
	case ApproachSiteCodes127:
		return "Lamina of fourth thoracic vertebra"
	case ApproachSiteCodes128:
		return "Intervertebral disc of eleventh thoracic vertebra"
	case ApproachSiteCodes129:
		return "Coccygeal plexus"
	case ApproachSiteCodes130:
		return "Nucleus pulposus of intervertebral disc of third lumbar vertebra"
	case ApproachSiteCodes131:
		return "Ventral lateral nucleus of thalamus"
	case ApproachSiteCodes132:
		return "Ileal arteries"
	case ApproachSiteCodes133:
		return "Entire symphysis"
	case ApproachSiteCodes134:
		return "Splenius capitis muscle"
	case ApproachSiteCodes135:
		return "Histioblast"
	case ApproachSiteCodes136:
		return "Otoconia"
	case ApproachSiteCodes137:
		return "Entire paramammary lymph node"
	case ApproachSiteCodes138:
		return "Intrinsic larynx"
	case ApproachSiteCodes139:
		return "Metaphase nucleus"
	case ApproachSiteCodes140:
		return "Third thoracic vertebra"
	case ApproachSiteCodes141:
		return "Medial collateral ligament of knee joint"
	case ApproachSiteCodes142:
		return "Supraorbital vein"
	case ApproachSiteCodes143:
		return "Foregut"
	case ApproachSiteCodes144:
		return "Hilum of left lung"
	case ApproachSiteCodes145:
		return "Transverse peduncular tract nucleus"
	case ApproachSiteCodes146:
		return "Dorsomedial nucleus of thalamus"
	case ApproachSiteCodes147:
		return "Ligamentum teres of liver"
	case ApproachSiteCodes148:
		return "Thymic lobule"
	case ApproachSiteCodes149:
		return "Ventral nuclear group of thalamus"
	case ApproachSiteCodes150:
		return "Periorbital region"
	case ApproachSiteCodes151:
		return "Cupula ampullaris"
	case ApproachSiteCodes152:
		return "Right tonsil"
	case ApproachSiteCodes153:
		return "Central tegmental tract"
	case ApproachSiteCodes154:
		return "Thoracic duct"
	case ApproachSiteCodes155:
		return "Lymphatic of thorax"
	case ApproachSiteCodes156:
		return "Premelanosome"
	case ApproachSiteCodes157:
		return "Sacroiliac region"
	case ApproachSiteCodes158:
		return "Naris"
	case ApproachSiteCodes159:
		return "Greater circulus arteriosus of iris"
	case ApproachSiteCodes160:
		return "Entire root of nose"
	case ApproachSiteCodes161:
		return "Bulbar conjunctiva"
	case ApproachSiteCodes162:
		return "Arrector pili muscles"
	case ApproachSiteCodes163:
		return "Pharyngeal recess"
	case ApproachSiteCodes164:
		return "Suprahyoid muscles"
	case ApproachSiteCodes165:
		return "Entire promontory common iliac lymph node"
	case ApproachSiteCodes166:
		return "Entire musculophrenic vein"
	case ApproachSiteCodes167:
		return "Skin of external ear"
	case ApproachSiteCodes168:
		return "Entire ear"
	case ApproachSiteCodes169:
		return "Suprarenal aorta"
	case ApproachSiteCodes170:
		return "Entire left elbow region"
	case ApproachSiteCodes171:
		return "Porus acusticus internus"
	case ApproachSiteCodes172:
		return "Cingulum dentis"
	case ApproachSiteCodes173:
		return "Clavicular facet of scapula"
	case ApproachSiteCodes174:
		return "Superior thoracic artery"
	case ApproachSiteCodes175:
		return "Spinal cord ventral median fissure"
	case ApproachSiteCodes176:
		return "Right fallopian tube"
	case ApproachSiteCodes177:
		return "Vaginal nerves"
	case ApproachSiteCodes178:
		return "Lingual tonsil"
	case ApproachSiteCodes179:
		return "Chorionic villi"
	case ApproachSiteCodes180:
		return "Skin of ear lobule"
	case ApproachSiteCodes181:
		return "Reticular formation of spinal cord"
	case ApproachSiteCodes182:
		return "Head of phalanx of hand"
	case ApproachSiteCodes183:
		return "Nucleus ambiguus"
	case ApproachSiteCodes184:
		return "Accessory sinus"
	case ApproachSiteCodes185:
		return "Tuberomammillary nucleus"
	case ApproachSiteCodes186:
		return "Urinary tract transitional epithelial cell"
	case ApproachSiteCodes187:
		return "Glial cell"
	case ApproachSiteCodes188:
		return "Ligamentum arteriosum"
	case ApproachSiteCodes189:
		return "Pharyngeal cavity"
	case ApproachSiteCodes190:
		return "Endometrial zona basalis"
	case ApproachSiteCodes191:
		return "Clavicular part of pectoralis major muscle"
	case ApproachSiteCodes192:
		return "Lamina of fifth thoracic vertebra"
	case ApproachSiteCodes193:
		return "Cerebral basal surface"
	case ApproachSiteCodes194:
		return "Lesser osseous pelvis"
	case ApproachSiteCodes195:
		return "Type I hair cell"
	case ApproachSiteCodes196:
		return "Subserosa"
	case ApproachSiteCodes197:
		return "Nasopharyngeal gland"
	case ApproachSiteCodes198:
		return "Veins of the knee"
	case ApproachSiteCodes199:
		return "Spinous process of cervical vertebra"
	case ApproachSiteCodes200:
		return "Base of third metacarpal bone"
	case ApproachSiteCodes201:
		return "Entire salivary seromucous gland (body structure)"
	case ApproachSiteCodes202:
		return "Segmental bronchial branches"
	case ApproachSiteCodes203:
		return "Metencephalon of fetus"
	case ApproachSiteCodes204:
		return "Calyx"
	case ApproachSiteCodes205:
		return "Nasal suture of skull"
	case ApproachSiteCodes206:
		return "Medial surface of toe"
	case ApproachSiteCodes207:
		return "Papillary muscles of right ventricle"
	case ApproachSiteCodes208:
		return "Superior margin of adrenal gland"
	case ApproachSiteCodes209:
		return "Transverse facial artery"
	case ApproachSiteCodes210:
		return "First metatarsal facet of medial cuneiform bone"
	case ApproachSiteCodes211:
		return "Mandibular left first premolar tooth"
	case ApproachSiteCodes212:
		return "Dorsum of foot"
	case ApproachSiteCodes213:
		return "Submaxillary ganglion"
	case ApproachSiteCodes214:
		return "Digital tendon and tendon sheath of foot"
	case ApproachSiteCodes215:
		return "Tunica intima of vein"
	case ApproachSiteCodes216:
		return "Subcutaneous tissue of posterior surface of forearm"
	case ApproachSiteCodes217:
		return "Articular surface, third metacarpal, of second metacarpal bone"
	case ApproachSiteCodes218:
		return "Skin of frenulum of clitoris"
	case ApproachSiteCodes219:
		return "Medial check ligament of eye"
	case ApproachSiteCodes220:
		return "Entire cisterna pontis"
	case ApproachSiteCodes221:
		return "Membrane of lysosome"
	case ApproachSiteCodes222:
		return "Pancreatic plexus"
	case ApproachSiteCodes223:
		return "Femoral triangle"
	case ApproachSiteCodes224:
		return "Superior rectal artery"
	case ApproachSiteCodes225:
		return "Cuboid articular facet of fourth metatarsal bone"
	case ApproachSiteCodes226:
		return "Phalanx of thumb"
	case ApproachSiteCodes227:
		return "Gracilis muscle"
	case ApproachSiteCodes228:
		return "Plasmablast"
	case ApproachSiteCodes229:
		return "All extremities"
	case ApproachSiteCodes230:
		return "Flexor pollicis longus muscle tendon"
	case ApproachSiteCodes231:
		return "Intervertebral disc of third thoracic vertebra"
	case ApproachSiteCodes232:
		return "Neuroendocrine tissue"
	case ApproachSiteCodes233:
		return "Posterior thalamic radiation of internal capsule"
	case ApproachSiteCodes234:
		return "Semispinalis capitis muscle"
	case ApproachSiteCodes235:
		return "Anterior cutaneous branch of lumbosacral plexus"
	case ApproachSiteCodes236:
		return "Anterior ethmoidal artery"
	case ApproachSiteCodes237:
		return "Dorsal nerve of penis"
	case ApproachSiteCodes238:
		return "Mucous membrane of urinary bladder"
	case ApproachSiteCodes239:
		return "Medial olfactory gyrus"
	case ApproachSiteCodes240:
		return "Bowman's space"
	case ApproachSiteCodes241:
		return "Left maxillary sinus"
	case ApproachSiteCodes242:
		return "Entire calcarine artery"
	case ApproachSiteCodes243:
		return "Capsule of ankle joint"
	case ApproachSiteCodes244:
		return "Apical foramen of tooth"
	case ApproachSiteCodes245:
		return "Fold for stapes"
	case ApproachSiteCodes246:
		return "Entire vitelline vein of placenta"
	case ApproachSiteCodes247:
		return "Endometrium"
	case ApproachSiteCodes248:
		return "Medial occipitotemporal gyrus"
	case ApproachSiteCodes249:
		return "Circular layer of gastric muscularis"
	case ApproachSiteCodes250:
		return "Spinal cord"
	case ApproachSiteCodes251:
		return "Eccrine gland"
	case ApproachSiteCodes252:
		return "Lamina propria of ureter"
	case ApproachSiteCodes253:
		return "Apocrine gland"
	case ApproachSiteCodes254:
		return "Pars tensa of tympanic membrane"
	case ApproachSiteCodes255:
		return "Tendon sheath of popliteus muscle"
	case ApproachSiteCodes256:
		return "Cremasteric fascia"
	case ApproachSiteCodes257:
		return "Head of femur"
	case ApproachSiteCodes258:
		return "Spinous process of fourth thoracic vertebra"
	case ApproachSiteCodes259:
		return "Lamina of fourth lumbar vertebra"
	case ApproachSiteCodes260:
		return "Dorsal digital nerves of lateral hallux and medial second toe"
	case ApproachSiteCodes261:
		return "Perivesicular tissues of seminal vesicles"
	case ApproachSiteCodes262:
		return "Renal artery"
	case ApproachSiteCodes263:
		return "Respiratory epithelium"
	case ApproachSiteCodes264:
		return "Superficial epigastric artery"
	case ApproachSiteCodes265:
		return "Accessory cephalic vein"
	case ApproachSiteCodes266:
		return "Entire gland (organ)"
	case ApproachSiteCodes267:
		return "Posterior epiglottis"
	case ApproachSiteCodes268:
		return "Anterior ligament of uterus"
	case ApproachSiteCodes269:
		return "Posterior portion of diaphragmatic aspect of liver"
	case ApproachSiteCodes270:
		return "Facial nerve motor branch"
	case ApproachSiteCodes271:
		return "Posterior papillary muscle of left ventricle"
	case ApproachSiteCodes272:
		return "Subcutaneous tissue of supraorbital area"
	case ApproachSiteCodes273:
		return "Anatomical space structure (body structure)"
	case ApproachSiteCodes274:
		return "Medial cuneiform bone"
	case ApproachSiteCodes275:
		return "Talar facet of navicular bone of foot"
	case ApproachSiteCodes276:
		return "Entire right margin of uterus"
	case ApproachSiteCodes277:
		return "Anterior limb of internal capsule"
	case ApproachSiteCodes278:
		return "White fibrocartilage"
	case ApproachSiteCodes279:
		return "Transitional epithelial cell"
	case ApproachSiteCodes280:
		return "Subcutaneous tissue of thigh"
	case ApproachSiteCodes281:
		return "Glomerular urinary pole"
	case ApproachSiteCodes282:
		return "Articular surface, metacarpal, of phalanx of thumb"
	case ApproachSiteCodes283:
		return "Bone marrow of vertebral body"
	case ApproachSiteCodes284:
		return "Anteroventral nucleus of thalamus"
	case ApproachSiteCodes285:
		return "Nerve"
	case ApproachSiteCodes286:
		return "Peripheral nervous system"
	case ApproachSiteCodes287:
		return "Spinal arachnoid"
	case ApproachSiteCodes288:
		return "Seminal vesicle lumen"
	case ApproachSiteCodes289:
		return "Mitochondrion in division"
	case ApproachSiteCodes290:
		return "Tendinous arch of pelvic fascia"
	case ApproachSiteCodes291:
		return "Clinical crown of tooth"
	case ApproachSiteCodes292:
		return "Suprachoroidal space"
	case ApproachSiteCodes293:
		return "Dorsal surface of index finger"
	case ApproachSiteCodes294:
		return "Trabecula carnea of left ventricle"
	case ApproachSiteCodes295:
		return "Pleura"
	case ApproachSiteCodes296:
		return "Head of fourth metatarsal bone"
	case ApproachSiteCodes297:
		return "Bony tissue"
	case ApproachSiteCodes298:
		return "Tractus olivocochlearis"
	case ApproachSiteCodes299:
		return "Obturator artery"
	case ApproachSiteCodes300:
		return "Costocervical trunk"
	case ApproachSiteCodes301:
		return "Spinal nerve"
	case ApproachSiteCodes302:
		return "Raphe of soft palate"
	case ApproachSiteCodes303:
		return "Endocardium of right atrium"
	case ApproachSiteCodes304:
		return "Monostomatic sublingual gland"
	case ApproachSiteCodes305:
		return "Subcutaneous tissue of nuchal region"
	case ApproachSiteCodes306:
		return "All large arteries"
	case ApproachSiteCodes307:
		return "Left coronary artery main stem"
	case ApproachSiteCodes308:
		return "Posterior segment of right upper lobe of lung"
	case ApproachSiteCodes309:
		return "Parametrial lymph node"
	case ApproachSiteCodes310:
		return "Papillary area"
	case ApproachSiteCodes311:
		return "Root canal of tooth"
	case ApproachSiteCodes312:
		return "Pedicle of third cervical vertebra"
	case ApproachSiteCodes313:
		return "Ventral anterior nucleus"
	case ApproachSiteCodes314:
		return "Tectopontine fibers"
	case ApproachSiteCodes315:
		return "Splenic branches of splenic artery"
	case ApproachSiteCodes316:
		return "Vestibulospinal tract"
	case ApproachSiteCodes317:
		return "Occipitofrontal diameter of head"
	case ApproachSiteCodes318:
		return "Haversian canal"
	case ApproachSiteCodes319:
		return "Right lung"
	case ApproachSiteCodes320:
		return "Entire right commissure of pulmonic valve"
	case ApproachSiteCodes321:
		return "Intertragal incisure"
	case ApproachSiteCodes322:
		return "Anterior papillary muscle of left ventricle"
	case ApproachSiteCodes323:
		return "Supporting tissue of rectum"
	case ApproachSiteCodes324:
		return "Secondary spermatocyte"
	case ApproachSiteCodes325:
		return "Agger nasi"
	case ApproachSiteCodes326:
		return "Rima oris"
	case ApproachSiteCodes327:
		return "Nonsegmented basophil"
	case ApproachSiteCodes328:
		return "Suboccipitobregmatic diameter of head"
	case ApproachSiteCodes329:
		return "Superior palpebral arch"
	case ApproachSiteCodes330:
		return "Mesogastrium"
	case ApproachSiteCodes331:
		return "Cell of bone"
	case ApproachSiteCodes332:
		return "Lateral margin of forearm"
	case ApproachSiteCodes333:
		return "Rotator muscles"
	case ApproachSiteCodes334:
		return "Deep lymphatics of upper extremity"
	case ApproachSiteCodes335:
		return "Thalamostriate veins"
	case ApproachSiteCodes336:
		return "Penetrated oocyte"
	case ApproachSiteCodes337:
		return "Anterodorsal nucleus of thalamus"
	case ApproachSiteCodes338:
		return "Commissure of tricuspid valve"
	case ApproachSiteCodes339:
		return "Posterior midline of trunk"
	case ApproachSiteCodes340:
		return "Vastus medialis muscle"
	case ApproachSiteCodes341:
		return "Embryonic testis"
	case ApproachSiteCodes342:
		return "Annulate lamella, cisternal lumen"
	case ApproachSiteCodes343:
		return "Subcutaneous tissue of suboccipital region"
	case ApproachSiteCodes344:
		return "Lateral wall of mastoid antrum"
	case ApproachSiteCodes345:
		return "Capsule of proximal tibiofibular joint"
	case ApproachSiteCodes346:
		return "Dorsal metatarsal artery"
	case ApproachSiteCodes347:
		return "Thyroid capsule"
	case ApproachSiteCodes348:
		return "Dorsal nucleus of trapezoid body"
	case ApproachSiteCodes349:
		return "Muscularis of ureter"
	case ApproachSiteCodes350:
		return "Body of vertebra"
	case ApproachSiteCodes351:
		return "Body of gallbladder"
	case ApproachSiteCodes352:
		return "Gastrophrenic ligament"
	case ApproachSiteCodes353:
		return "Arch of tenth thoracic vertebra"
	case ApproachSiteCodes354:
		return "Straight part of longus colli muscle"
	case ApproachSiteCodes355:
		return "Ischiococcygeus muscle"
	case ApproachSiteCodes356:
		return "Occipital branch of posterior auricular artery"
	case ApproachSiteCodes357:
		return "Lamellipodium"
	case ApproachSiteCodes358:
		return "Tympanic ostium of Eustachian tube"
	case ApproachSiteCodes359:
		return "Pelvic wall"
	case ApproachSiteCodes360:
		return "Entire subpyloric lymph node"
	case ApproachSiteCodes361:
		return "Great vessel"
	case ApproachSiteCodes362:
		return "Lateral thoracic artery"
	case ApproachSiteCodes363:
		return "Nucleus pulposus of intervertebral disc of first thoracic vertebra"
	case ApproachSiteCodes364:
		return "Subcutaneous tissue of lower extremity"
	case ApproachSiteCodes365:
		return "Entire dorsal metacarpal ligament"
	case ApproachSiteCodes366:
		return "Superior segment of right lower lobe of lung"
	case ApproachSiteCodes367:
		return "Argentaffin cells"
	case ApproachSiteCodes368:
		return "Nasal septal cartilage"
	case ApproachSiteCodes369:
		return "Apex of coccyx"
	case ApproachSiteCodes370:
		return "Transplanted liver"
	case ApproachSiteCodes371:
		return "Interscapular region of back"
	case ApproachSiteCodes372:
		return "Dorsal surface of great toe"
	case ApproachSiteCodes373:
		return "Subcutaneous tissue of femoral region"
	case ApproachSiteCodes374:
		return "Common carotid plexus"
	case ApproachSiteCodes375:
		return "Subcutaneous tissue of lateral surface of fourth toe"
	case ApproachSiteCodes376:
		return "Occipital lymph node"
	case ApproachSiteCodes377:
		return "Pericardiophrenic artery"
	case ApproachSiteCodes378:
		return "Vestibular window"
	case ApproachSiteCodes379:
		return "Head of tenth rib"
	case ApproachSiteCodes380:
		return "Entorhinal cortex"
	case ApproachSiteCodes381:
		return "Lacrimal sac"
	case ApproachSiteCodes382:
		return "Fifth metatarsal articular facet of fourth metatarsal bone"
	case ApproachSiteCodes383:
		return "Rectus capitis muscle"
	case ApproachSiteCodes384:
		return "Olfactory tract"
	case ApproachSiteCodes385:
		return "Gyrus of brain"
	case ApproachSiteCodes386:
		return "Entire parietal branch of anterior cerebral artery"
	case ApproachSiteCodes387:
		return "Subcutaneous tissue of concha"
	case ApproachSiteCodes388:
		return "Deep veins of clitoris"
	case ApproachSiteCodes389:
		return "Medial nuclei of globus pallidus"
	case ApproachSiteCodes390:
		return "Chromosomes, group A"
	case ApproachSiteCodes391:
		return "Posterior commissure of labia majora"
	case ApproachSiteCodes392:
		return "Eosinophilic promyelocyte"
	case ApproachSiteCodes393:
		return "Lateral wall of orbit"
	case ApproachSiteCodes394:
		return "Capsule of proximal interphalangeal joint of index finger"
	case ApproachSiteCodes395:
		return "Fourth coccygeal vertebra"
	case ApproachSiteCodes396:
		return "Entire dorsal lingual vein"
	case ApproachSiteCodes397:
		return "Vagus nerve bronchial branches"
	case ApproachSiteCodes398:
		return "Macula of saccule"
	case ApproachSiteCodes399:
		return "Lumbosacral spinal cord central canal"
	case ApproachSiteCodes400:
		return "Superior frontal sulcus"
	case ApproachSiteCodes401:
		return "Artery of extremity"
	case ApproachSiteCodes402:
		return "Palmar surface of little finger"
	case ApproachSiteCodes403:
		return "Celiac plexus"
	case ApproachSiteCodes404:
		return "Abdominal aortic plexus"
	case ApproachSiteCodes405:
		return "Hyparterial bronchi"
	case ApproachSiteCodes406:
		return "Both lower extremities"
	case ApproachSiteCodes407:
		return "Entire extensor tendon and tendon sheath of fifth toe"
	case ApproachSiteCodes408:
		return "Türk cell"
	case ApproachSiteCodes409:
		return "Epithelial cell"
	case ApproachSiteCodes410:
		return "Head of second rib"
	case ApproachSiteCodes411:
		return "First metacarpal bone"
	case ApproachSiteCodes412:
		return "Posterior tibial veins"
	case ApproachSiteCodes413:
		return "Inferior articular process of seventh cervical vertebra"
	case ApproachSiteCodes414:
		return "Middle portion of ileum"
	case ApproachSiteCodes415:
		return "Paracortical area of lymph node"
	case ApproachSiteCodes416:
		return "Cartilage canal"
	case ApproachSiteCodes417:
		return "Anterior midline of abdomen"
	case ApproachSiteCodes418:
		return "Spinalis muscle"
	case ApproachSiteCodes419:
		return "Protoplasmic astrocyte"
	case ApproachSiteCodes420:
		return "Upper jaw"
	case ApproachSiteCodes421:
		return "Subchorionic space"
	case ApproachSiteCodes422:
		return "Lateral surface of little finger"
	case ApproachSiteCodes423:
		return "Stratum spinosum"
	case ApproachSiteCodes424:
		return "Small intestine mucous membrane"
	case ApproachSiteCodes425:
		return "Fourth metatarsal facet of lateral cuneiform bone"
	case ApproachSiteCodes426:
		return "Incisure of cartilaginous portion of auditory canal"
	case ApproachSiteCodes427:
		return "Parafascicular nucleus of thalamus"
	case ApproachSiteCodes428:
		return "Scala vestibuli"
	case ApproachSiteCodes429:
		return "Anterior articular surface for talus"
	case ApproachSiteCodes430:
		return "Tracheal submucosa"
	case ApproachSiteCodes431:
		return "Cell"
	case ApproachSiteCodes432:
		return "Clivus ossis sphenoidalis"
	case ApproachSiteCodes433:
		return "Ductus arteriosus"
	case ApproachSiteCodes434:
		return "Dental arch"
	case ApproachSiteCodes435:
		return "Accessory sinus gland"
	case ApproachSiteCodes436:
		return "Subclavian plexus"
	case ApproachSiteCodes437:
		return "Joint of lower extremity"
	case ApproachSiteCodes438:
		return "Internal medullary lamina of thalamus"
	case ApproachSiteCodes439:
		return "Lamellated granule, as in surfactant-secreting cell"
	case ApproachSiteCodes440:
		return "Entire vagus nerve parasympathetic fibers to liver, gallbladder, bile ducts and pancreas"
	case ApproachSiteCodes441:
		return "Tentorium cerebelli"
	case ApproachSiteCodes442:
		return "Desmosome"
	case ApproachSiteCodes443:
		return "Skin of posterior surface of thigh"
	case ApproachSiteCodes444:
		return "Middle trunk of brachial plexus"
	case ApproachSiteCodes445:
		return "Larynx"
	case ApproachSiteCodes446:
		return "Base of phalanx of foot"
	case ApproachSiteCodes447:
		return "Tubercle of eighth rib"
	case ApproachSiteCodes448:
		return "Lesser tuberosity of humerus"
	case ApproachSiteCodes449:
		return "Medullary cord"
	case ApproachSiteCodes450:
		return "Lipid droplet, homogeneous"
	case ApproachSiteCodes451:
		return "Tunica albuginea of corpus spongiosum"
	case ApproachSiteCodes452:
		return "Skin of nuchal region"
	case ApproachSiteCodes453:
		return "Basal lamina, inclusion in subepithelial location"
	case ApproachSiteCodes454:
		return "Cardinal vein"
	case ApproachSiteCodes455:
		return "Neutrophilic myelocyte"
	case ApproachSiteCodes456:
		return "Entire venous plexus of the foramen ovale basis cranii"
	case ApproachSiteCodes457:
		return "Ventral sacrococcygeal ligament"
	case ApproachSiteCodes458:
		return "Medial surface of great toe"
	case ApproachSiteCodes459:
		return "Gemellus muscle"
	case ApproachSiteCodes460:
		return "Supracardinal vein"
	case ApproachSiteCodes461:
		return "Perineal nerve"
	case ApproachSiteCodes462:
		return "Phrenic nerve pericardial branch"
	case ApproachSiteCodes463:
		return "Ventral posterior inferior nucleus"
	case ApproachSiteCodes464:
		return "Deiter's cell"
	case ApproachSiteCodes465:
		return "Uterine venous plexus"
	case ApproachSiteCodes466:
		return "Anterior tibial compartment"
	case ApproachSiteCodes467:
		return "Femoral canal"
	case ApproachSiteCodes468:
		return "Subcutaneous tissue of ring finger"
	case ApproachSiteCodes469:
		return "Membranous ampulla"
	case ApproachSiteCodes470:
		return "Tuberculum impar"
	case ApproachSiteCodes471:
		return "Constrictor pharyngeus muscle"
	case ApproachSiteCodes472:
		return "Dorsal tegmental nuclei of midbrain"
	case ApproachSiteCodes473:
		return "Spiral lamina of modiolus"
	case ApproachSiteCodes474:
		return "Entire sublingual vein"
	case ApproachSiteCodes475:
		return "Entire interlobular vein of kidney"
	case ApproachSiteCodes476:
		return "Cell membrane, prokaryotic"
	case ApproachSiteCodes477:
		return "Uterovaginal plexus"
	case ApproachSiteCodes478:
		return "Tympanic antrum"
	case ApproachSiteCodes479:
		return "Cerebellar gracile lobule"
	case ApproachSiteCodes480:
		return "Lymph node of lower extremity"
	case ApproachSiteCodes481:
		return "Radial notch of ulna"
	case ApproachSiteCodes482:
		return "Subcutaneous tissue of back"
	case ApproachSiteCodes483:
		return "Amygdaloid nucleus"
	case ApproachSiteCodes484:
		return "Superior temporal sulcus"
	case ApproachSiteCodes485:
		return "Yellow bone marrow"
	case ApproachSiteCodes486:
		return "Posterior surface of prostate"
	case ApproachSiteCodes487:
		return "Superficial dorsal veins of clitoris"
	case ApproachSiteCodes488:
		return "Obturator internus muscle ischial bursa"
	case ApproachSiteCodes489:
		return "Rugal column"
	case ApproachSiteCodes490:
		return "Infrasternal angle"
	case ApproachSiteCodes491:
		return "Posterior auricular vein"
	case ApproachSiteCodes492:
		return "Entire angle of first rib"
	case ApproachSiteCodes493:
		return "Suspensory ligament of lens"
	case ApproachSiteCodes494:
		return "Intervertebral foramen of twelfth thoracic vertebra"
	case ApproachSiteCodes495:
		return "Epithelium of lens"
	case ApproachSiteCodes496:
		return "Right external carotid artery"
	case ApproachSiteCodes497:
		return "Superior ileocecal recess"
	case ApproachSiteCodes498:
		return "Frontal veins"
	case ApproachSiteCodes499:
		return "Uterine ostium of fallopian tube"
	case ApproachSiteCodes500:
		return "Right cerebral hemisphere"
	case ApproachSiteCodes501:
		return "Mucosa of gallbladder"
	case ApproachSiteCodes502:
		return "Intervertebral disc of thoracic vertebra"
	case ApproachSiteCodes503:
		return "Skin of lateral portion of neck"
	case ApproachSiteCodes504:
		return "Foramen singulare"
	case ApproachSiteCodes505:
		return "Anterior mediastinal lymph node"
	case ApproachSiteCodes506:
		return "Superior pole of kidney"
	case ApproachSiteCodes507:
		return "Fourth cervical vertebra"
	case ApproachSiteCodes508:
		return "Inferior frontal gyrus"
	case ApproachSiteCodes509:
		return "Synaptic specialization, cytoplasmic"
	case ApproachSiteCodes510:
		return "Median arcuate ligament of diaphragm"
	case ApproachSiteCodes511:
		return "Hippocampus"
	case ApproachSiteCodes512:
		return "Small intestine muscularis propria"
	case ApproachSiteCodes513:
		return "Superior fascia of perineum"
	case ApproachSiteCodes514:
		return "Uterine paracervical lymph node"
	case ApproachSiteCodes515:
		return "Normal fat pad"
	case ApproachSiteCodes516:
		return "Articular process of third lumbar vertebra"
	case ApproachSiteCodes517:
		return "Sex chromosome Y"
	case ApproachSiteCodes518:
		return "Apocrine intraepidermal duct"
	case ApproachSiteCodes519:
		return "Deep artery of clitoris"
	case ApproachSiteCodes520:
		return "Cardiac incisure of stomach"
	case ApproachSiteCodes521:
		return "Lacrimal part of orbicularis oculi muscle"
	case ApproachSiteCodes522:
		return "Metacarpophalangeal joint of little finger"
	case ApproachSiteCodes523:
		return "Superior aberrant ductule of epididymis"
	case ApproachSiteCodes524:
		return "Subcutaneous tissue of chin"
	case ApproachSiteCodes525:
		return "Tegmental portion of pons"
	case ApproachSiteCodes526:
		return "Longitudinal layer of duodenal muscularis propria"
	case ApproachSiteCodes527:
		return "Alveolar ridge mucous membrane"
	case ApproachSiteCodes528:
		return "Singlet"
	case ApproachSiteCodes529:
		return "Seventh costal cartilage"
	case ApproachSiteCodes530:
		return "Tendon of supraspinatus muscle"
	case ApproachSiteCodes531:
		return "Retina of right eye"
	case ApproachSiteCodes532:
		return "Anulus fibrosus of intervertebral disc of fifth cervical vertebra"
	case ApproachSiteCodes533:
		return "Navicular facet of intermediate cuneiform bone"
	case ApproachSiteCodes534:
		return "Right visceral pleura"
	case ApproachSiteCodes535:
		return "Muscular portion of interventricular septum"
	case ApproachSiteCodes536:
		return "Canal of stomach"
	case ApproachSiteCodes537:
		return "Fractured membrane P face"
	case ApproachSiteCodes538:
		return "Entire inner surface of seventh rib"
	case ApproachSiteCodes539:
		return "Retina"
	case ApproachSiteCodes540:
		return "Lower digestive tract"
	case ApproachSiteCodes541:
		return "Subcutaneous tissue of upper extremity"
	case ApproachSiteCodes542:
		return "Entire articular part of tubercle of ninth rib"
	case ApproachSiteCodes543:
		return "Skin of lateral surface of finger"
	case ApproachSiteCodes544:
		return "Multifidus muscles"
	case ApproachSiteCodes545:
		return "Submandibular triangle"
	case ApproachSiteCodes546:
		return "Temporal fossa"
	case ApproachSiteCodes547:
		return "Tendon and tendon sheath of leg and ankle"
	case ApproachSiteCodes548:
		return "Anterior cervical lymph node"
	case ApproachSiteCodes549:
		return "Skin of forearm"
	case ApproachSiteCodes550:
		return "Subcutaneous tissue of anterior portion of neck"
	case ApproachSiteCodes551:
		return "Endocervical epithelium"
	case ApproachSiteCodes552:
		return "Paradidymis"
	case ApproachSiteCodes553:
		return "Diaphragm"
	case ApproachSiteCodes554:
		return "Medium sized neuron"
	case ApproachSiteCodes555:
		return "Entire angle of seventh rib"
	case ApproachSiteCodes556:
		return "Superior rectus muscle"
	case ApproachSiteCodes557:
		return "Duodenal fold"
	case ApproachSiteCodes558:
		return "Substantia propria of sclera"
	case ApproachSiteCodes559:
		return "Posterior cord of brachial plexus"
	case ApproachSiteCodes560:
		return "Superior articular process of seventh cervical vertebra"
	case ApproachSiteCodes561:
		return "Orbital plate of ethmoid bone"
	case ApproachSiteCodes562:
		return "Serosa of urinary bladder"
	case ApproachSiteCodes563:
		return "Subcutaneous tissue of lateral border of sole of foot"
	case ApproachSiteCodes564:
		return "Tuberosity of distal phalanx of hand"
	case ApproachSiteCodes565:
		return "Endothelial sieve plate"
	case ApproachSiteCodes566:
		return "Articular surface, third metacarpal, of fourth metacarpal bone"
	case ApproachSiteCodes567:
		return "Posterior cells of ethmoid sinus"
	case ApproachSiteCodes568:
		return "Superior recess of tympanic membrane"
	case ApproachSiteCodes569:
		return "Myotome"
	case ApproachSiteCodes570:
		return "Articular process of twelfth thoracic vertebra"
	case ApproachSiteCodes571:
		return "Bronchial lumen"
	case ApproachSiteCodes572:
		return "Great cardiac vein"
	case ApproachSiteCodes573:
		return "Tensor tympani muscle"
	case ApproachSiteCodes574:
		return "Vestibular vein"
	case ApproachSiteCodes575:
		return "Posterior palatine arch"
	case ApproachSiteCodes576:
		return "Capsule of distal interphalangeal joint of third toe"
	case ApproachSiteCodes577:
		return "Left wrist"
	case ApproachSiteCodes578:
		return "Eighth rib"
	case ApproachSiteCodes579:
		return "Subcutaneous tissue of eyelid"
	case ApproachSiteCodes580:
		return "Episcleral artery"
	case ApproachSiteCodes581:
		return "Chromosomes, group D"
	case ApproachSiteCodes582:
		return "Quadratus lumborum muscle"
	case ApproachSiteCodes583:
		return "Intervertebral disc of second thoracic vertebra"
	case ApproachSiteCodes584:
		return "Circular layer of duodenal muscularis propria"
	case ApproachSiteCodes585:
		return "Mesentery of ascending colon"
	case ApproachSiteCodes586:
		return "Penicilliary arteries"
	case ApproachSiteCodes587:
		return "Heterolysosome"
	case ApproachSiteCodes588:
		return "Columnar epithelial cell"
	case ApproachSiteCodes589:
		return "Entire outer surface of third rib"
	case ApproachSiteCodes590:
		return "Lacrimal vein"
	case ApproachSiteCodes591:
		return "Metacarpophalangeal joint of middle finger"
	case ApproachSiteCodes592:
		return "Deciduous mandibular right canine tooth"
	case ApproachSiteCodes593:
		return "Ligament of left superior vena cava"
	case ApproachSiteCodes594:
		return "Capsule of temporomandibular joint"
	case ApproachSiteCodes595:
		return "Gastrointestinal subserosa"
	case ApproachSiteCodes596:
		return "Subclavian nerve"
	case ApproachSiteCodes597:
		return "Body of fifth thoracic vertebra"
	case ApproachSiteCodes598:
		return "Facial nerve parasympathetic fibers"
	case ApproachSiteCodes599:
		return "Entire postcapillary venule"
	case ApproachSiteCodes600:
		return "Piriform recess"
	case ApproachSiteCodes601:
		return "Os lacrimale"
	case ApproachSiteCodes602:
		return "Sulcus terminalis cordis"
	case ApproachSiteCodes603:
		return "Accessory phrenic nerves"
	case ApproachSiteCodes604:
		return "Subcutaneous tissue of scalp"
	case ApproachSiteCodes605:
		return "Skin of dorsal surface of finger"
	case ApproachSiteCodes606:
		return "Posterior basal branch of left pulmonary artery"
	case ApproachSiteCodes607:
		return "Aryepiglottic muscle"
	case ApproachSiteCodes608:
		return "Structure of fetal atloid articulation"
	case ApproachSiteCodes609:
		return "Lymphoid follicle of stomach"
	case ApproachSiteCodes610:
		return "Hair medulla"
	case ApproachSiteCodes611:
		return "Lymphatics of thyroid gland"
	case ApproachSiteCodes612:
		return "Cavernous portion of urethra"
	case ApproachSiteCodes613:
		return "Coccygeal nerve"
	case ApproachSiteCodes614:
		return "Ligamentum nuchae"
	case ApproachSiteCodes615:
		return "Presymphysial lymph node"
	case ApproachSiteCodes616:
		return "Medial malleolus"
	case ApproachSiteCodes617:
		return "Supraspinatus muscle"
	case ApproachSiteCodes618:
		return "Structure of radiating portion of cortical lobule of kidney"
	case ApproachSiteCodes619:
		return "Mast cell"
	case ApproachSiteCodes620:
		return "Posterior vagal trunk"
	case ApproachSiteCodes621:
		return "Cytotrophoblast"
	case ApproachSiteCodes622:
		return "Medial aspect of ovary"
	case ApproachSiteCodes623:
		return "Glans clitoridis"
	case ApproachSiteCodes624:
		return "Distal portion of circumflex branch of left coronary artery"
	case ApproachSiteCodes625:
		return "Cardiac valve leaflet"
	case ApproachSiteCodes626:
		return "Colonic haustra"
	case ApproachSiteCodes627:
		return "Thyrocervical trunk"
	case ApproachSiteCodes628:
		return "Entire anterior commissure of mitral valve"
	case ApproachSiteCodes629:
		return "Gastrohepatic ligament"
	case ApproachSiteCodes630:
		return "Angular incisure of stomach"
	case ApproachSiteCodes631:
		return "Pollicis artery"
	case ApproachSiteCodes632:
		return "Inferior nasal turbinate"
	case ApproachSiteCodes633:
		return "Medial border of sole"
	case ApproachSiteCodes634:
		return "Cerebellar hemisphere"
	case ApproachSiteCodes635:
		return "Base of phalanx of middle finger"
	case ApproachSiteCodes636:
		return "Lingual nerve"
	case ApproachSiteCodes637:
		return "Structure of dorsal intercuneiform ligaments"
	case ApproachSiteCodes638:
		return "Sphenoparietal sinus"
	case ApproachSiteCodes639:
		return "Cuticle of nail"
	case ApproachSiteCodes640:
		return "Sternal muscle"
	case ApproachSiteCodes641:
		return "Entire right posterior cerebral artery"
	case ApproachSiteCodes642:
		return "Entire right anterior cerebral artery"
	case ApproachSiteCodes643:
		return "Anterior fossa of cranial cavity"
	case ApproachSiteCodes644:
		return "Uterine subserosa"
	case ApproachSiteCodes645:
		return "Central lobule of cerebellum"
	case ApproachSiteCodes646:
		return "Articular facet of head of fibula"
	case ApproachSiteCodes647:
		return "Right ankle"
	case ApproachSiteCodes648:
		return "Arch of second lumbar vertebra"
	case ApproachSiteCodes649:
		return "Femoral nerve lateral muscular branches"
	case ApproachSiteCodes650:
		return "Pleural recess"
	case ApproachSiteCodes651:
		return "Chorda tympani"
	case ApproachSiteCodes652:
		return "Entire callosomarginal branch of anterior cerebral artery"
	case ApproachSiteCodes653:
		return "Mitochondrial inclusion"
	case ApproachSiteCodes654:
		return "Right knee"
	case ApproachSiteCodes655:
		return "Structure of tendon and tendon sheath of hand"
	case ApproachSiteCodes656:
		return "Spermatozoa"
	case ApproachSiteCodes657:
		return "Macula of utricle"
	case ApproachSiteCodes658:
		return "Entire interstitial tissue of spleen"
	case ApproachSiteCodes659:
		return "Obturator nerve anterior branch"
	case ApproachSiteCodes660:
		return "Ligament of lumbosacral joint"
	case ApproachSiteCodes661:
		return "Pars ciliaris of retina"
	case ApproachSiteCodes662:
		return "Axial skeleton"
	case ApproachSiteCodes663:
		return "Corticomedullary junction of kidney"
	case ApproachSiteCodes664:
		return "Spore crystal"
	case ApproachSiteCodes665:
		return "Secondary foot process"
	case ApproachSiteCodes666:
		return "Leaf of epiglottis"
	case ApproachSiteCodes667:
		return "Habenular commissure"
	case ApproachSiteCodes668:
		return "Visceral pericardium"
	case ApproachSiteCodes669:
		return "Medial surface of arm"
	case ApproachSiteCodes670:
		return "Popliteal region"
	case ApproachSiteCodes671:
		return "Subcutaneous tissue of medial surface of third toe"
	case ApproachSiteCodes672:
		return "Lower alveolar ridge mucosa"
	case ApproachSiteCodes673:
		return "Perivascular space"
	case ApproachSiteCodes674:
		return "Right upper extremity"
	case ApproachSiteCodes675:
		return "Entire jugular arch"
	case ApproachSiteCodes676:
		return "Entire anterior labial vein"
	case ApproachSiteCodes677:
		return "Lymphocytic tissue"
	case ApproachSiteCodes678:
		return "Anterior myocardium"
	case ApproachSiteCodes679:
		return "Posterior hypothalamic nucleus"
	case ApproachSiteCodes680:
		return "Collateral sulcus"
	case ApproachSiteCodes681:
		return "Thoracolumbar region of back"
	case ApproachSiteCodes682:
		return "Subcutaneous tissue of jaw"
	case ApproachSiteCodes683:
		return "Bile duct mucous membrane"
	case ApproachSiteCodes684:
		return "Subcutaneous tissue of external genitalia"
	case ApproachSiteCodes685:
		return "Right colic artery"
	case ApproachSiteCodes686:
		return "Interstitial tissue of myocardium"
	case ApproachSiteCodes687:
		return "Middle phalanx of index finger"
	case ApproachSiteCodes688:
		return "Ventral posterolateral nucleus of thalamus"
	case ApproachSiteCodes689:
		return "Attachment plaque of desmosome or hemidesmosome"
	case ApproachSiteCodes690:
		return "Fetal implantation site"
	case ApproachSiteCodes691:
		return "Anulus fibrosus of intervertebral disc of thoracic vertebra"
	case ApproachSiteCodes692:
		return "False rib"
	case ApproachSiteCodes693:
		return "Trigeminal ganglion sensory root"
	case ApproachSiteCodes694:
		return "Base of metacarpal bone"
	case ApproachSiteCodes695:
		return "Paraduodenal recess"
	case ApproachSiteCodes696:
		return "Cauda equina"
	case ApproachSiteCodes697:
		return "Gustatory pore"
	case ApproachSiteCodes698:
		return "Isthmus tympani posticus"
	case ApproachSiteCodes699:
		return "Hypoglossal nerve intrinsic tongue muscle branch"
	case ApproachSiteCodes700:
		return "Entire inferior choroid vein"
	case ApproachSiteCodes701:
		return "Appendiceal muscularis propria"
	case ApproachSiteCodes702:
		return "Muscle of perineum"
	case ApproachSiteCodes703:
		return "Deep inguinal ring"
	case ApproachSiteCodes704:
		return "Anterior surface of arm"
	case ApproachSiteCodes705:
		return "Lingual gyrus"
	case ApproachSiteCodes706:
		return "Ciliary processes"
	case ApproachSiteCodes707:
		return "Lymphatic of head"
	case ApproachSiteCodes708:
		return "Entire left margin of uterus"
	case ApproachSiteCodes709:
		return "Paraventricular nucleus of thalamus"
	case ApproachSiteCodes710:
		return "Entire plantar calcaneocuboidal ligament"
	case ApproachSiteCodes711:
		return "Anterior semicircular duct"
	case ApproachSiteCodes712:
		return "Ovarian ligament"
	case ApproachSiteCodes713:
		return "Lateral surface of sublingual gland"
	case ApproachSiteCodes714:
		return "Lipid, crystalline"
	case ApproachSiteCodes715:
		return "Iliotibial tract"
	case ApproachSiteCodes716:
		return "Cerebellar lenticular nucleus"
	case ApproachSiteCodes717:
		return "Entire plantar tarsal ligament"
	case ApproachSiteCodes718:
		return "Entire anterior ligament of head of fibula"
	case ApproachSiteCodes719:
		return "Entire vasa vasorum"
	case ApproachSiteCodes720:
		return "Vagus nerve parasympathetic fibers"
	case ApproachSiteCodes721:
		return "Deep head of flexor pollicis brevis muscle"
	case ApproachSiteCodes722:
		return "Mitotic cell in anaphase"
	case ApproachSiteCodes723:
		return "Finger"
	case ApproachSiteCodes724:
		return "Intervertebral disc space of eleventh thoracic vertebra"
	case ApproachSiteCodes725:
		return "Subcutaneous tissue of vertex"
	case ApproachSiteCodes726:
		return "Connexon"
	case ApproachSiteCodes727:
		return "Tenth thoracic vertebra"
	case ApproachSiteCodes728:
		return "Thalamoolivary tract"
	case ApproachSiteCodes729:
		return "Intervenous tubercle of right atrium"
	case ApproachSiteCodes730:
		return "Frenulum labii"
	case ApproachSiteCodes731:
		return "Femoral artery"
	case ApproachSiteCodes732:
		return "Subtendinous bursa of triceps brachii muscle"
	case ApproachSiteCodes733:
		return "Pontine portion of medial longitudinal fasciculus"
	case ApproachSiteCodes734:
		return "Subdural space of spinal region"
	case ApproachSiteCodes735:
		return "Skin of medial surface of fifth toe"
	case ApproachSiteCodes736:
		return "Posterior choroidal artery"
	case ApproachSiteCodes737:
		return "Palatine duct"
	case ApproachSiteCodes738:
		return "Skin appendage"
	case ApproachSiteCodes739:
		return "Mesovarian margin of ovary"
	case ApproachSiteCodes740:
		return "Lamina of third thoracic vertebra"
	case ApproachSiteCodes741:
		return "Entire striate artery"
	case ApproachSiteCodes742:
		return "Right foot"
	case ApproachSiteCodes743:
		return "Sympathetic trunk spinal nerve branch"
	case ApproachSiteCodes744:
		return "Lateral posterior nucleus of thalamus"
	case ApproachSiteCodes745:
		return "Anterior surface of manubrium"
	case ApproachSiteCodes746:
		return "Abdominal aorta"
	case ApproachSiteCodes747:
		return "Posterior margin of nasal septum"
	case ApproachSiteCodes748:
		return "Subcutaneous tissue of submental area"
	case ApproachSiteCodes749:
		return "Macrocytic normochromic erythrocyte"
	case ApproachSiteCodes750:
		return "Sternoclavicular joint"
	case ApproachSiteCodes751:
		return "Intracranial subdural space"
	case ApproachSiteCodes752:
		return "Mandibular canal"
	case ApproachSiteCodes753:
		return "Myocardium of ventricle"
	case ApproachSiteCodes754:
		return "Scapular region of back"
	case ApproachSiteCodes755:
		return "Rhopheocytotic vesicle"
	case ApproachSiteCodes756:
		return "Corneal corpuscle"
	case ApproachSiteCodes757:
		return "Rotator cuff including muscles and tendons"
	case ApproachSiteCodes758:
		return "Submucosa of anal canal"
	case ApproachSiteCodes759:
		return "Occipital angle of parietal bone"
	case ApproachSiteCodes760:
		return "Olivocerebellar fibers"
	case ApproachSiteCodes761:
		return "Proximal phalanx of third toe"
	case ApproachSiteCodes762:
		return "Ligament of diaphragm"
	case ApproachSiteCodes763:
		return "Helper cell"
	case ApproachSiteCodes764:
		return "Lamina propria of ethmoid sinus"
	case ApproachSiteCodes765:
		return "Entire first left aortic arch"
	case ApproachSiteCodes766:
		return "Abdominopelvic portion of sympathetic nervous system"
	case ApproachSiteCodes767:
		return "Skin of glans penis"
	case ApproachSiteCodes768:
		return "Articulations of auditory ossicles"
	case ApproachSiteCodes769:
		return "Mucous membrane of tongue"
	case ApproachSiteCodes770:
		return "Anterior communicating artery"
	case ApproachSiteCodes771:
		return "Inflow tract of right ventricle"
	case ApproachSiteCodes772:
		return "Limitans nucleus"
	case ApproachSiteCodes773:
		return "Subcutaneous acromial bursa"
	case ApproachSiteCodes774:
		return "Superficial flexor tendon of little finger"
	case ApproachSiteCodes775:
		return "Membrane-coating granule, amorphous"
	case ApproachSiteCodes776:
		return "Lateral nuclei of globus pallidus"
	case ApproachSiteCodes777:
		return "Pancreatic veins"
	case ApproachSiteCodes778:
		return "Entire superficial circumflex iliac vein"
	case ApproachSiteCodes779:
		return "Stratum lemnisci of corpora quadrigemina"
	case ApproachSiteCodes780:
		return "Radial nerve"
	case ApproachSiteCodes781:
		return "Intervertebral disc space of twelfth thoracic vertebra"
	case ApproachSiteCodes782:
		return "Infundibulum of Fallopian tube"
	case ApproachSiteCodes783:
		return "Intranuclear crystal"
	case ApproachSiteCodes784:
		return "Hindgut"
	case ApproachSiteCodes785:
		return "Entire delphian lymph node"
	case ApproachSiteCodes786:
		return "Supraaortic valve area"
	case ApproachSiteCodes787:
		return "Superior anastomotic vein"
	case ApproachSiteCodes788:
		return "Vein of head"
	case ApproachSiteCodes789:
		return "Interlobar duct of pancreas"
	case ApproachSiteCodes790:
		return "Superior colliculus of corpora quadrigemina"
	case ApproachSiteCodes791:
		return "Entire lateral striate artery"
	case ApproachSiteCodes792:
		return "Infraorbital nerve"
	case ApproachSiteCodes793:
		return "Superior articular process of fifth thoracic vertebra"
	case ApproachSiteCodes794:
		return "Wrist"
	case ApproachSiteCodes795:
		return "Accessory atrioventricular bundle"
	case ApproachSiteCodes796:
		return "Apical branch of right pulmonary artery"
	case ApproachSiteCodes797:
		return "Osseous portion of Eustachian tube"
	case ApproachSiteCodes798:
		return "Tunica interna of eyeball"
	case ApproachSiteCodes799:
		return "Articular surface, metacarpal, of phalanx of hand"
	case ApproachSiteCodes800:
		return "Small intestine serosa"
	case ApproachSiteCodes801:
		return "Below knee region"
	case ApproachSiteCodes802:
		return "Interlobular arteries of liver"
	case ApproachSiteCodes803:
		return "Mastoid fontanel of skull"
	case ApproachSiteCodes804:
		return "Lumbar lymph node"
	case ApproachSiteCodes805:
		return "Colic lymph node"
	case ApproachSiteCodes806:
		return "Sphincter pupillae muscle"
	case ApproachSiteCodes807:
		return "Jugum of sphenoid bone"
	case ApproachSiteCodes808:
		return "Lamina of eighth thoracic vertebra"
	case ApproachSiteCodes809:
		return "Birth canal"
	case ApproachSiteCodes810:
		return "Iliac fossa"
	case ApproachSiteCodes811:
		return "Renal surface of adrenal gland"
	case ApproachSiteCodes812:
		return "Joint of lumbar vertebra"
	case ApproachSiteCodes813:
		return "Ligament of sacroiliac joint and pubic symphysis"
	case ApproachSiteCodes814:
		return "Sinoatrial node branch of right coronary artery"
	case ApproachSiteCodes815:
		return "Mesial surface of tooth"
	case ApproachSiteCodes816:
		return "Obliquus capitis muscle"
	case ApproachSiteCodes817:
		return "Inferior articular process of twelfth thoracic vertebra"
	case ApproachSiteCodes818:
		return "Posterior intercavernous sinus"
	case ApproachSiteCodes819:
		return "Lipid droplet"
	case ApproachSiteCodes820:
		return "Entire juxtaintestinal lymph node"
	case ApproachSiteCodes821:
		return "Interclavicular ligament"
	case ApproachSiteCodes822:
		return "Both feet"
	case ApproachSiteCodes823:
		return "Meissner's plexus"
	case ApproachSiteCodes824:
		return "Vestibulocochlear nerve"
	case ApproachSiteCodes825:
		return "Cricoid cartilage"
	case ApproachSiteCodes826:
		return "Adductor hallucis muscle"
	case ApproachSiteCodes827:
		return "Medulla oblongata fasciculus cuneatus"
	case ApproachSiteCodes828:
		return "Right margin of heart"
	case ApproachSiteCodes829:
		return "Zygomatic region of face"
	case ApproachSiteCodes830:
		return "Transplanted ureter"
	case ApproachSiteCodes831:
		return "Superior right pulmonary vein"
	case ApproachSiteCodes832:
		return "Choroidal branches of posterior spinal artery"
	case ApproachSiteCodes833:
		return "Glycogen vacuole"
	case ApproachSiteCodes834:
		return "All toes"
	case ApproachSiteCodes835:
		return "Body of right atrium"
	case ApproachSiteCodes836:
		return "Lateral olfactory gyrus"
	case ApproachSiteCodes837:
		return "Intervertebral foramen of second lumbar vertebra"
	case ApproachSiteCodes838:
		return "Entire minor sublingual ducts"
	case ApproachSiteCodes839:
		return "Periodontal tissues"
	case ApproachSiteCodes840:
		return "Subcutaneous tissue of interdigital space of hand"
	case ApproachSiteCodes841:
		return "Cavernous portion of internal carotid artery"
	case ApproachSiteCodes842:
		return "Tendinous arch"
	case ApproachSiteCodes843:
		return "Intranuclear body, granular with filamentous capsule"
	case ApproachSiteCodes844:
		return "Corticomedullary junction of adrenal gland"
	case ApproachSiteCodes845:
		return "Iliac tuberosity"
	case ApproachSiteCodes846:
		return "Thenar and hypothenar spaces"
	case ApproachSiteCodes847:
		return "Pedicle of eleventh thoracic vertebra"
	case ApproachSiteCodes848:
		return "Peroneal artery"
	case ApproachSiteCodes849:
		return "Shaft of phalanx of middle finger"
	case ApproachSiteCodes850:
		return "Agranular endoplasmic reticulum, connection with other organelle"
	case ApproachSiteCodes851:
		return "Entire subtendinous prepatellar bursa"
	case ApproachSiteCodes852:
		return "Proper fasciculus"
	case ApproachSiteCodes853:
		return "Crista galli"
	case ApproachSiteCodes854:
		return "Palmar surface of middle finger"
	case ApproachSiteCodes855:
		return "Mandibular right second premolar tooth"
	case ApproachSiteCodes856:
		return "Brachiocephalic vein"
	case ApproachSiteCodes857:
		return "Diaphragmatic surface of lung"
	case ApproachSiteCodes858:
		return "Entire gastric cardiac gland (body structure)"
	case ApproachSiteCodes859:
		return "Lateral glossoepiglottic fold"
	case ApproachSiteCodes860:
		return "Entire left ulnar artery"
	case ApproachSiteCodes861:
		return "Entire inferior transverse scapular ligament"
	case ApproachSiteCodes862:
		return "Entire endocardium of right ventricle"
	case ApproachSiteCodes863:
		return "Inguinal lymph node"
	case ApproachSiteCodes864:
		return "Coracoid process of scapula"
	case ApproachSiteCodes865:
		return "Cerebral meninges"
	case ApproachSiteCodes866:
		return "Trapezoid ligament"
	case ApproachSiteCodes867:
		return "Stratum zonale of corpora quadrigemina"
	case ApproachSiteCodes868:
		return "Left eye"
	case ApproachSiteCodes869:
		return "Joint of vertebral column"
	case ApproachSiteCodes870:
		return "Marginal part of orbicularis oris muscle"
	case ApproachSiteCodes871:
		return "Hepatic vein"
	case ApproachSiteCodes872:
		return "Cerebellar peduncle"
	case ApproachSiteCodes873:
		return "Left parietal lobe"
	case ApproachSiteCodes874:
		return "Entire middle colic vein"
	case ApproachSiteCodes875:
		return "Ascending colon"
	case ApproachSiteCodes876:
		return "Both forearms"
	case ApproachSiteCodes877:
		return "White matter of insula"
	case ApproachSiteCodes878:
		return "Splenic sinusoid"
	case ApproachSiteCodes879:
		return "Superior laryngeal vein"
	case ApproachSiteCodes880:
		return "Arch of foot"
	case ApproachSiteCodes881:
		return "Vein of the scala tympani"
	case ApproachSiteCodes882:
		return "Transverse folds of palate"
	case ApproachSiteCodes883:
		return "Embryo stage 1 structure"
	case ApproachSiteCodes884:
		return "Capsule of metatarsophalangeal joint of fifth toe"
	case ApproachSiteCodes885:
		return "Filaments of contractile apparatus"
	case ApproachSiteCodes886:
		return "Intervertebral disc of eighth thoracic vertebra"
	case ApproachSiteCodes887:
		return "Centriole"
	case ApproachSiteCodes888:
		return "Shaft of fifth metatarsal bone"
	case ApproachSiteCodes889:
		return "Structure of lumbar rotator muscle (body structure)"
	case ApproachSiteCodes890:
		return "Entire external pudendal vein"
	case ApproachSiteCodes891:
		return "Niemann-Pick cell"
	case ApproachSiteCodes892:
		return "Posterior segment of right lobe of liver"
	case ApproachSiteCodes893:
		return "Gravid uterus"
	case ApproachSiteCodes894:
		return "Tendon and tendon sheath of second toe"
	case ApproachSiteCodes895:
		return "Pelvic fascia"
	case ApproachSiteCodes896:
		return "Corpus cavernosum of penis"
	case ApproachSiteCodes897:
		return "Posterior intraoccipital synchondrosis"
	case ApproachSiteCodes898:
		return "Labial veins"
	case ApproachSiteCodes899:
		return "Merkel's tactile disc"
	case ApproachSiteCodes900:
		return "Subtendinous iliac bursa"
	case ApproachSiteCodes901:
		return "Tail of epididymis"
	case ApproachSiteCodes902:
		return "Interdental papilla of gingiva"
	case ApproachSiteCodes903:
		return "Lateral ligament of temporomandibular joint"
	case ApproachSiteCodes904:
		return "Skin of medial surface of middle finger"
	case ApproachSiteCodes905:
		return "All permanent teeth"
	case ApproachSiteCodes906:
		return "Pecten ani"
	case ApproachSiteCodes907:
		return "Lumbar vein"
	case ApproachSiteCodes908:
		return "Lymphatics of stomach"
	case ApproachSiteCodes909:
		return "Plantar surface of fourth toe"
	case ApproachSiteCodes910:
		return "Structure of deep cervical lymphatics"
	case ApproachSiteCodes911:
		return "Subclavian vein"
	case ApproachSiteCodes912:
		return "Structure of medial cartilaginous lamina of pharyngotympanic tube (body structure)"
	case ApproachSiteCodes913:
		return "Structure of amacrine cell of retina (body structure)"
	case ApproachSiteCodes914:
		return "Afferent glomerular arteriole"
	case ApproachSiteCodes915:
		return "Pulmonary ligament"
	case ApproachSiteCodes916:
		return "Head of metacarpal bone"
	case ApproachSiteCodes917:
		return "Coronal depression of tooth"
	case ApproachSiteCodes918:
		return "Calcaneocuboidal ligament"
	case ApproachSiteCodes919:
		return "Pyramid of medulla oblongata"
	case ApproachSiteCodes920:
		return "Entire facet for fifth costal cartilage of sternum"
	case ApproachSiteCodes921:
		return "Duodenal lumen"
	case ApproachSiteCodes922:
		return "Subcutaneous tissue of areola"
	case ApproachSiteCodes923:
		return "Deep branch of ulnar nerve"
	case ApproachSiteCodes924:
		return "Posterior process of nasal septal cartilage"
	case ApproachSiteCodes925:
		return "Lanugo hair"
	case ApproachSiteCodes926:
		return "Left superior vena cava"
	case ApproachSiteCodes927:
		return "Entire superior transverse scapular ligament"
	case ApproachSiteCodes928:
		return "Gastric mucous gland"
	case ApproachSiteCodes929:
		return "Infraclavicular lymph node"
	case ApproachSiteCodes930:
		return "Subcutaneous tissue of lower margin of nasal septum"
	case ApproachSiteCodes931:
		return "Ciliary muscle"
	case ApproachSiteCodes932:
		return "Head of second metatarsal bone"
	case ApproachSiteCodes933:
		return "Melanocyte"
	case ApproachSiteCodes934:
		return "Posterior scrotal branches of internal pudendal artery"
	case ApproachSiteCodes935:
		return "Iliac fascia"
	case ApproachSiteCodes936:
		return "Right wrist"
	case ApproachSiteCodes937:
		return "Entire tendon of index finger"
	case ApproachSiteCodes938:
		return "Submucosa of tonsil"
	case ApproachSiteCodes939:
		return "Genital tubercle"
	case ApproachSiteCodes940:
		return "Left carotid sinus"
	case ApproachSiteCodes941:
		return "Distinctive shape of mitochondrial cristae"
	case ApproachSiteCodes942:
		return "Superficial lymphatics of thorax"
	case ApproachSiteCodes943:
		return "Deep venous system of lower extremity"
	case ApproachSiteCodes944:
		return "Skeletal muscle fiber, type IIb"
	case ApproachSiteCodes945:
		return "Fascia of upper extremity"
	case ApproachSiteCodes946:
		return "Proximal phalanx of little toe"
	case ApproachSiteCodes947:
		return "Perforating branches of internal thoracic artery"
	case ApproachSiteCodes948:
		return "Biparietal diameter of head"
	case ApproachSiteCodes949:
		return "Interspinalis thoracis muscles"
	case ApproachSiteCodes950:
		return "Right kidney"
	case ApproachSiteCodes951:
		return "Hilum of adrenal gland"
	case ApproachSiteCodes952:
		return "Fornix of lacrimal sac"
	case ApproachSiteCodes953:
		return "Carunculae hymenales"
	case ApproachSiteCodes954:
		return "Thymus"
	case ApproachSiteCodes955:
		return "Entire appendicular vein"
	case ApproachSiteCodes956:
		return "Thyroid tubercle"
	case ApproachSiteCodes957:
		return "Peripheral nerve myelinated nerve fiber"
	case ApproachSiteCodes958:
		return "Transverse arytenoid muscle"
	case ApproachSiteCodes959:
		return "Paracentral lobule"
	case ApproachSiteCodes960:
		return "Posterior ethmoidal nerve"
	case ApproachSiteCodes961:
		return "Primary foot process"
	case ApproachSiteCodes962:
		return "Ileocecal ostium"
	case ApproachSiteCodes963:
		return "Rhomboideus cervicis muscle"
	case ApproachSiteCodes964:
		return "Superior articular process of sixth thoracic vertebra"
	case ApproachSiteCodes965:
		return "Duodenal ampulla"
	case ApproachSiteCodes966:
		return "Lateral meniscus of knee joint"
	case ApproachSiteCodes967:
		return "Base of lung"
	case ApproachSiteCodes968:
		return "Base of phalanx of index finger"
	case ApproachSiteCodes969:
		return "Ventral spinocerebellar tract of pons"
	case ApproachSiteCodes970:
		return "Nucleus pulposus of intervertebral disc of eighth thoracic vertebra"
	case ApproachSiteCodes971:
		return "Intervertebral foramen of fifth thoracic vertebra"
	case ApproachSiteCodes972:
		return "Transplanted lung"
	case ApproachSiteCodes973:
		return "Male"
	case ApproachSiteCodes974:
		return "Ophthalmic nerve"
	case ApproachSiteCodes975:
		return "Levator labii superioris muscle"
	case ApproachSiteCodes976:
		return "Deep volar arch of radial artery"
	case ApproachSiteCodes977:
		return "Deep dorsal sacrococcygeal ligament"
	case ApproachSiteCodes978:
		return "Medial surface of third toe"
	case ApproachSiteCodes979:
		return "Zygomatic nerve"
	case ApproachSiteCodes980:
		return "Vagus nerve pharyngeal plexus"
	case ApproachSiteCodes981:
		return "Entire costal groove of fifth rib"
	case ApproachSiteCodes982:
		return "Merkel's cell"
	case ApproachSiteCodes983:
		return "Middle part of cartilaginous portion of auditory canal"
	case ApproachSiteCodes984:
		return "Entire inner surface of eleventh rib"
	case ApproachSiteCodes985:
		return "Supraoptic region of hypothalamus"
	case ApproachSiteCodes986:
		return "Liver"
	case ApproachSiteCodes987:
		return "Gynecoid pelvis"
	case ApproachSiteCodes988:
		return "Parotid lymph node"
	case ApproachSiteCodes989:
		return "Abductor digiti minimi muscle of foot"
	case ApproachSiteCodes990:
		return "Entire anterior commissure of aortic valve"
	case ApproachSiteCodes991:
		return "Iliac artery"
	case ApproachSiteCodes992:
		return "Uropod"
	case ApproachSiteCodes993:
		return "Entire anteromedial perforating artery"
	case ApproachSiteCodes994:
		return "Capsule of calcaneocuboidal joint"
	case ApproachSiteCodes995:
		return "Anulus fibrosus of intervertebral disc of tenth thoracic vertebra"
	case ApproachSiteCodes996:
		return "Sweat gland"
	case ApproachSiteCodes997:
		return "Pigmented layer of ciliary body"
	case ApproachSiteCodes998:
		return "Lesser sciatic notch"
	case ApproachSiteCodes999:
		return "Pulmonic area"
	default:
		return "Unknown Action Participant Type"
	}
}

/*
Anatomical structure
Posterior carpal region
Fetal part of placenta
Entire condylar emissary vein
Visceral layer of Bowman's capsule
Parathyroid gland
Subcutaneous tissue of medial surface of index finger
Coronoid process of mandible
Central pair of microtubules, cilium or flagellum, not bacterial
Deep circumflex artery of ilium
Supraclavicular part of brachial plexus
Anterior division of renal artery
Entire left commissure of aortic valve
Gluteus maximus muscle
Articular surface, phalanges, of fourth metacarpal bone
Canal of Hering
Hepatocolic ligament
Superior labial artery
Lateral vestibular nucleus
Mesotympanum
Pectoral region
Kupffer cell
Thoracic nerve
Right lower lobe of lung
Superior articular process of lumbar vertebra
Lateral myocardium
Central axillary lymph node
Entire flexor tendon and tendon sheath of fourth toe
Metacarpophalangeal joint of index finger
Fifth metatarsal bone
Plantar surface of great toe
Skin of umbilicus
Cardiac impression of liver
Ankle
Atrioventricular bundle
Reticular corium
Wall of urinary bladder
Dental branches of inferior alveolar artery
Entire posterior temporal diploic vein
Gastric fundus
Inferior surface of tongue
Trochanteric bursa
Collateral ligament
Lateral corticospinal tract
Basophilic normoblast
Ascending frontal gyrus
Flexor hallucis longus muscle
Cardiopulmonary circulatory system
Transverse colon
Costal surface of lung
Entire vagus nerve parasympathetic fibers to cardiac plexus
Intervertebral disc space of fifth lumbar vertebra
Head of phalanx of great toe
Capsule of proximal interphalangeal joint of third toe
Interventricular septum
Palpebral fissure
Subcutaneous tissue of philtrum
Multivesicular body, internal vesicles
Tuberosity of distal phalanx of little toe
Superior articular process of seventh thoracic vertebra
Tracheal mucous membrane
Jaw
Embryonic structure
Fetal hyaloid artery
Small intestine submucosa
Body of ischium
Dense intermediate filament bundles
Head and neck structure
Visceral aspect of liver
Entire deep temporal vein
Posterior intercostal arteries
Fetal chondrocranium
Posterior cervical spinal cord nerve root
Spinous process of fifth thoracic vertebra
Oral region of face
Lamina muscularis of colonic mucous membrane
Anterior cruciate ligament of knee joint
Superior laryngeal aperture
Thyrohyoid branch of ansa cervicalis
Crus of diaphragm
Bronchus
Ovarian vein
Meningeal branches of occipital artery
Entire diaphragmatic lymph node
Fibrous portion of pericardium
Peritonsillar tissue
Sebaceous gland
Vesicular bursa of sternohyoid muscle
Frontozygomatic suture of skull
Promonocyte
Entire subcutaneous prepatellar bursa
Female
Sternothyroid muscle
Superior occipital gyrus
Thymic cortex
Cranial cavity
Entire major calyx
Tarsal gland
Inferior longitudinal muscle of tongue
Aortopulmonary septum
Frenulum linguae
Odontoid process of axis
Mandibular nerve
Chromosomes, group E
Teres major muscle
Synostosis
Meninges
Duodenal serosa
Inferior articular process of sixth cervical vertebra
Dorsal digital nerves of radial nerve
Distinctive arrangement of microtubules
Vertebral nerve
Glottis
Telogen hair
Deep flexor tendon of index finger
Gastric serosa
Vastus lateralis muscle
Posterior limb of stapes
Paravesicular lymph node
Ventricle of larynx
Yellow fibrocartilage
Parietal branch of superficial temporal artery
Metatarsus
Soft tissues of trunk
Anterior cecal artery
Ejaculatory duct
Frontomental diameter of head
Lamina of fourth thoracic vertebra
Intervertebral disc of eleventh thoracic vertebra
Coccygeal plexus
Nucleus pulposus of intervertebral disc of third lumbar vertebra
Ventral lateral nucleus of thalamus
Ileal arteries
Entire symphysis
Splenius capitis muscle
Histioblast
Otoconia
Entire paramammary lymph node
Intrinsic larynx
Metaphase nucleus
Third thoracic vertebra
Medial collateral ligament of knee joint
Supraorbital vein
Foregut
Hilum of left lung
Transverse peduncular tract nucleus
Dorsomedial nucleus of thalamus
Ligamentum teres of liver
Thymic lobule
Ventral nuclear group of thalamus
Periorbital region
Cupula ampullaris
Right tonsil
Central tegmental tract
Thoracic duct
Lymphatic of thorax
Premelanosome
Sacroiliac region
Naris
Greater circulus arteriosus of iris
Entire root of nose
Bulbar conjunctiva
Arrector pili muscles
Pharyngeal recess
Suprahyoid muscles
Entire promontory common iliac lymph node
Entire musculophrenic vein
Skin of external ear
Entire ear
Suprarenal aorta
Entire left elbow region
Porus acusticus internus
Cingulum dentis
Clavicular facet of scapula
Superior thoracic artery
Spinal cord ventral median fissure
Right fallopian tube
Vaginal nerves
Lingual tonsil
Chorionic villi
Skin of ear lobule
Reticular formation of spinal cord
Head of phalanx of hand
Nucleus ambiguus
Accessory sinus
Tuberomammillary nucleus
Urinary tract transitional epithelial cell
Glial cell
Ligamentum arteriosum
Pharyngeal cavity
Endometrial zona basalis
Clavicular part of pectoralis major muscle
Lamina of fifth thoracic vertebra
Cerebral basal surface
Lesser osseous pelvis
Type I hair cell
Subserosa
Nasopharyngeal gland
Veins of the knee
Spinous process of cervical vertebra
Base of third metacarpal bone
Entire salivary seromucous gland (body structure)
Segmental bronchial branches
Metencephalon of fetus
Calyx
Nasal suture of skull
Medial surface of toe
Papillary muscles of right ventricle
Superior margin of adrenal gland
Transverse facial artery
First metatarsal facet of medial cuneiform bone
Mandibular left first premolar tooth
Dorsum of foot
Submaxillary ganglion
Digital tendon and tendon sheath of foot
Tunica intima of vein
Subcutaneous tissue of posterior surface of forearm
Articular surface, third metacarpal, of second metacarpal bone
Skin of frenulum of clitoris
Medial check ligament of eye
Entire cisterna pontis
Membrane of lysosome
Pancreatic plexus
Femoral triangle
Superior rectal artery
Cuboid articular facet of fourth metatarsal bone
Phalanx of thumb
Gracilis muscle
Plasmablast
All extremities
Flexor pollicis longus muscle tendon
Intervertebral disc of third thoracic vertebra
Neuroendocrine tissue
Posterior thalamic radiation of internal capsule
Semispinalis capitis muscle
Anterior cutaneous branch of lumbosacral plexus
Anterior ethmoidal artery
Dorsal nerve of penis
Mucous membrane of urinary bladder
Medial olfactory gyrus
Bowman's space
Left maxillary sinus
Entire calcarine artery
Capsule of ankle joint
Apical foramen of tooth
Fold for stapes
Entire vitelline vein of placenta
Endometrium
Medial occipitotemporal gyrus
Circular layer of gastric muscularis
Spinal cord
Eccrine gland
Lamina propria of ureter
Apocrine gland
Pars tensa of tympanic membrane
Tendon sheath of popliteus muscle
Cremasteric fascia
Head of femur
Spinous process of fourth thoracic vertebra
Lamina of fourth lumbar vertebra
Dorsal digital nerves of lateral hallux and medial second toe
Perivesicular tissues of seminal vesicles
Renal artery
Respiratory epithelium
Superficial epigastric artery
Accessory cephalic vein
Entire gland (organ)
Posterior epiglottis
Anterior ligament of uterus
Posterior portion of diaphragmatic aspect of liver
Facial nerve motor branch
Posterior papillary muscle of left ventricle
Subcutaneous tissue of supraorbital area
Anatomical space structure (body structure)
Medial cuneiform bone
Talar facet of navicular bone of foot
Entire right margin of uterus
Anterior limb of internal capsule
White fibrocartilage
Transitional epithelial cell
Subcutaneous tissue of thigh
Glomerular urinary pole
Articular surface, metacarpal, of phalanx of thumb
Bone marrow of vertebral body
Anteroventral nucleus of thalamus
Nerve
Peripheral nervous system
Spinal arachnoid
Seminal vesicle lumen
Mitochondrion in division
Tendinous arch of pelvic fascia
Clinical crown of tooth
Suprachoroidal space
Dorsal surface of index finger
Trabecula carnea of left ventricle
Pleura
Head of fourth metatarsal bone
Bony tissue
Tractus olivocochlearis
Obturator artery
Costocervical trunk
Spinal nerve
Raphe of soft palate
Endocardium of right atrium
Monostomatic sublingual gland
Subcutaneous tissue of nuchal region
All large arteries
Left coronary artery main stem
Posterior segment of right upper lobe of lung
Parametrial lymph node
Papillary area
Root canal of tooth
Pedicle of third cervical vertebra
Ventral anterior nucleus
Tectopontine fibers
Splenic branches of splenic artery
Vestibulospinal tract
Occipitofrontal diameter of head
Haversian canal
Right lung
Entire right commissure of pulmonic valve
Intertragal incisure
Anterior papillary muscle of left ventricle
Supporting tissue of rectum
Secondary spermatocyte
Agger nasi
Rima oris
Nonsegmented basophil
Suboccipitobregmatic diameter of head
Superior palpebral arch
Mesogastrium
Cell of bone
Lateral margin of forearm
Rotator muscles
Deep lymphatics of upper extremity
Thalamostriate veins
Penetrated oocyte
Anterodorsal nucleus of thalamus
Commissure of tricuspid valve
Posterior midline of trunk
Vastus medialis muscle
Embryonic testis
Annulate lamella, cisternal lumen
Subcutaneous tissue of suboccipital region
Lateral wall of mastoid antrum
Capsule of proximal tibiofibular joint
Dorsal metatarsal artery
Thyroid capsule
Dorsal nucleus of trapezoid body
Muscularis of ureter
Body of vertebra
Body of gallbladder
Gastrophrenic ligament
Arch of tenth thoracic vertebra
Straight part of longus colli muscle
Ischiococcygeus muscle
Occipital branch of posterior auricular artery
Lamellipodium
Tympanic ostium of Eustachian tube
Pelvic wall
Entire subpyloric lymph node
Great vessel
Lateral thoracic artery
Nucleus pulposus of intervertebral disc of first thoracic vertebra
Subcutaneous tissue of lower extremity
Entire dorsal metacarpal ligament
Superior segment of right lower lobe of lung
Argentaffin cells
Nasal septal cartilage
Apex of coccyx
Transplanted liver
Interscapular region of back
Dorsal surface of great toe
Subcutaneous tissue of femoral region
Common carotid plexus
Subcutaneous tissue of lateral surface of fourth toe
Occipital lymph node
Pericardiophrenic artery
Vestibular window
Head of tenth rib
Entorhinal cortex
Lacrimal sac
Fifth metatarsal articular facet of fourth metatarsal bone
Rectus capitis muscle
Olfactory tract
Gyrus of brain
Entire parietal branch of anterior cerebral artery
Subcutaneous tissue of concha
Deep veins of clitoris
Medial nuclei of globus pallidus
Chromosomes, group A
Posterior commissure of labia majora
Eosinophilic promyelocyte
Lateral wall of orbit
Capsule of proximal interphalangeal joint of index finger
Fourth coccygeal vertebra
Entire dorsal lingual vein
Vagus nerve bronchial branches
Macula of saccule
Lumbosacral spinal cord central canal
Superior frontal sulcus
Artery of extremity
Palmar surface of little finger
Celiac plexus
Abdominal aortic plexus
Hyparterial bronchi
Both lower extremities
Entire extensor tendon and tendon sheath of fifth toe
Türk cell
Epithelial cell
Head of second rib
First metacarpal bone
Posterior tibial veins
Inferior articular process of seventh cervical vertebra
Middle portion of ileum
Paracortical area of lymph node
Cartilage canal
Anterior midline of abdomen
Spinalis muscle
Protoplasmic astrocyte
Upper jaw
Subchorionic space
Lateral surface of little finger
Stratum spinosum
Small intestine mucous membrane
Fourth metatarsal facet of lateral cuneiform bone
Incisure of cartilaginous portion of auditory canal
Parafascicular nucleus of thalamus
Scala vestibuli
Anterior articular surface for talus
Tracheal submucosa
Cell
Clivus ossis sphenoidalis
Ductus arteriosus
Dental arch
Accessory sinus gland
Subclavian plexus
Joint of lower extremity
Internal medullary lamina of thalamus
Lamellated granule, as in surfactant-secreting cell
Entire vagus nerve parasympathetic fibers to liver, gallbladder, bile ducts and pancreas
Tentorium cerebelli
Desmosome
Skin of posterior surface of thigh
Middle trunk of brachial plexus
Larynx
Base of phalanx of foot
Tubercle of eighth rib
Lesser tuberosity of humerus
Medullary cord
Lipid droplet, homogeneous
Tunica albuginea of corpus spongiosum
Skin of nuchal region
Basal lamina, inclusion in subepithelial location
Cardinal vein
Neutrophilic myelocyte
Entire venous plexus of the foramen ovale basis cranii
Ventral sacrococcygeal ligament
Medial surface of great toe
Gemellus muscle
Supracardinal vein
Perineal nerve
Phrenic nerve pericardial branch
Ventral posterior inferior nucleus
Deiter's cell
Uterine venous plexus
Anterior tibial compartment
Femoral canal
Subcutaneous tissue of ring finger
Membranous ampulla
Tuberculum impar
Constrictor pharyngeus muscle
Dorsal tegmental nuclei of midbrain
Spiral lamina of modiolus
Entire sublingual vein
Entire interlobular vein of kidney
Cell membrane, prokaryotic
Uterovaginal plexus
Tympanic antrum
Cerebellar gracile lobule
Lymph node of lower extremity
Radial notch of ulna
Subcutaneous tissue of back
Amygdaloid nucleus
Superior temporal sulcus
Yellow bone marrow
Posterior surface of prostate
Superficial dorsal veins of clitoris
Obturator internus muscle ischial bursa
Rugal column
Infrasternal angle
Posterior auricular vein
Entire angle of first rib
Suspensory ligament of lens
Intervertebral foramen of twelfth thoracic vertebra
Epithelium of lens
Right external carotid artery
Superior ileocecal recess
Frontal veins
Uterine ostium of fallopian tube
Right cerebral hemisphere
Mucosa of gallbladder
Intervertebral disc of thoracic vertebra
Skin of lateral portion of neck
Foramen singulare
Anterior mediastinal lymph node
Superior pole of kidney
Fourth cervical vertebra
Inferior frontal gyrus
Synaptic specialization, cytoplasmic
Median arcuate ligament of diaphragm
Hippocampus
Small intestine muscularis propria
Superior fascia of perineum
Uterine paracervical lymph node
Normal fat pad
Articular process of third lumbar vertebra
Sex chromosome Y
Apocrine intraepidermal duct
Deep artery of clitoris
Cardiac incisure of stomach
Lacrimal part of orbicularis oculi muscle
Metacarpophalangeal joint of little finger
Superior aberrant ductule of epididymis
Subcutaneous tissue of chin
Tegmental portion of pons
Longitudinal layer of duodenal muscularis propria
Alveolar ridge mucous membrane
Singlet
Seventh costal cartilage
Tendon of supraspinatus muscle
Retina of right eye
Anulus fibrosus of intervertebral disc of fifth cervical vertebra
Navicular facet of intermediate cuneiform bone
Right visceral pleura
Muscular portion of interventricular septum
Canal of stomach
Fractured membrane P face
Entire inner surface of seventh rib
Retina
Lower digestive tract
Subcutaneous tissue of upper extremity
Entire articular part of tubercle of ninth rib
Skin of lateral surface of finger
Multifidus muscles
Submandibular triangle
Temporal fossa
Tendon and tendon sheath of leg and ankle
Anterior cervical lymph node
Skin of forearm
Subcutaneous tissue of anterior portion of neck
Endocervical epithelium
Paradidymis
Diaphragm
Medium sized neuron
Entire angle of seventh rib
Superior rectus muscle
Duodenal fold
Substantia propria of sclera
Posterior cord of brachial plexus
Superior articular process of seventh cervical vertebra
Orbital plate of ethmoid bone
Serosa of urinary bladder
Subcutaneous tissue of lateral border of sole of foot
Tuberosity of distal phalanx of hand
Endothelial sieve plate
Articular surface, third metacarpal, of fourth metacarpal bone
Posterior cells of ethmoid sinus
Superior recess of tympanic membrane
Myotome
Articular process of twelfth thoracic vertebra
Bronchial lumen
Great cardiac vein
Tensor tympani muscle
Vestibular vein
Posterior palatine arch
Capsule of distal interphalangeal joint of third toe
Left wrist
Eighth rib
Subcutaneous tissue of eyelid
Episcleral artery
Chromosomes, group D
Quadratus lumborum muscle
Intervertebral disc of second thoracic vertebra
Circular layer of duodenal muscularis propria
Mesentery of ascending colon
Penicilliary arteries
Heterolysosome
Columnar epithelial cell
Entire outer surface of third rib
Lacrimal vein
Metacarpophalangeal joint of middle finger
Deciduous mandibular right canine tooth
Ligament of left superior vena cava
Capsule of temporomandibular joint
Gastrointestinal subserosa
Subclavian nerve
Body of fifth thoracic vertebra
Facial nerve parasympathetic fibers
Entire postcapillary venule
Piriform recess
Os lacrimale
Sulcus terminalis cordis
Accessory phrenic nerves
Subcutaneous tissue of scalp
Skin of dorsal surface of finger
Posterior basal branch of left pulmonary artery
Aryepiglottic muscle
Structure of fetal atloid articulation
Lymphoid follicle of stomach
Hair medulla
Lymphatics of thyroid gland
Cavernous portion of urethra
Coccygeal nerve
Ligamentum nuchae
Presymphysial lymph node
Medial malleolus
Supraspinatus muscle
Structure of radiating portion of cortical lobule of kidney
Mast cell
Posterior vagal trunk
Cytotrophoblast
Medial aspect of ovary
Glans clitoridis
Distal portion of circumflex branch of left coronary artery
Cardiac valve leaflet
Colonic haustra
Thyrocervical trunk
Entire anterior commissure of mitral valve
Gastrohepatic ligament
Angular incisure of stomach
Pollicis artery
Inferior nasal turbinate
Medial border of sole
Cerebellar hemisphere
Base of phalanx of middle finger
Lingual nerve
Structure of dorsal intercuneiform ligaments
Sphenoparietal sinus
Cuticle of nail
Sternal muscle
Entire right posterior cerebral artery
Entire right anterior cerebral artery
Anterior fossa of cranial cavity
Uterine subserosa
Central lobule of cerebellum
Articular facet of head of fibula
Right ankle
Arch of second lumbar vertebra
Femoral nerve lateral muscular branches
Pleural recess
Chorda tympani
Entire callosomarginal branch of anterior cerebral artery
Mitochondrial inclusion
Right knee
Structure of tendon and tendon sheath of hand
Spermatozoa
Macula of utricle
Entire interstitial tissue of spleen
Obturator nerve anterior branch
Ligament of lumbosacral joint
Pars ciliaris of retina
Axial skeleton
Corticomedullary junction of kidney
Spore crystal
Secondary foot process
Leaf of epiglottis
Habenular commissure
Visceral pericardium
Medial surface of arm
Popliteal region
Subcutaneous tissue of medial surface of third toe
Lower alveolar ridge mucosa
Perivascular space
Right upper extremity
Entire jugular arch
Entire anterior labial vein
Lymphocytic tissue
Anterior myocardium
Posterior hypothalamic nucleus
Collateral sulcus
Thoracolumbar region of back
Subcutaneous tissue of jaw
Bile duct mucous membrane
Subcutaneous tissue of external genitalia
Right colic artery
Interstitial tissue of myocardium
Middle phalanx of index finger
Ventral posterolateral nucleus of thalamus
Attachment plaque of desmosome or hemidesmosome
Fetal implantation site
Anulus fibrosus of intervertebral disc of thoracic vertebra
False rib
Trigeminal ganglion sensory root
Base of metacarpal bone
Paraduodenal recess
Cauda equina
Gustatory pore
Isthmus tympani posticus
Hypoglossal nerve intrinsic tongue muscle branch
Entire inferior choroid vein
Appendiceal muscularis propria
Muscle of perineum
Deep inguinal ring
Anterior surface of arm
Lingual gyrus
Ciliary processes
Lymphatic of head
Entire left margin of uterus
Paraventricular nucleus of thalamus
Entire plantar calcaneocuboidal ligament
Anterior semicircular duct
Ovarian ligament
Lateral surface of sublingual gland
Lipid, crystalline
Iliotibial tract
Cerebellar lenticular nucleus
Entire plantar tarsal ligament
Entire anterior ligament of head of fibula
Entire vasa vasorum
Vagus nerve parasympathetic fibers
Deep head of flexor pollicis brevis muscle
Mitotic cell in anaphase
Finger
Intervertebral disc space of eleventh thoracic vertebra
Subcutaneous tissue of vertex
Connexon
Tenth thoracic vertebra
Thalamoolivary tract
Intervenous tubercle of right atrium
Frenulum labii
Femoral artery
Subtendinous bursa of triceps brachii muscle
Pontine portion of medial longitudinal fasciculus
Subdural space of spinal region
Skin of medial surface of fifth toe
Posterior choroidal artery
Palatine duct
Skin appendage
Mesovarian margin of ovary
Lamina of third thoracic vertebra
Entire striate artery
Right foot
Sympathetic trunk spinal nerve branch
Lateral posterior nucleus of thalamus
Anterior surface of manubrium
Abdominal aorta
Posterior margin of nasal septum
Subcutaneous tissue of submental area
Macrocytic normochromic erythrocyte
Sternoclavicular joint
Intracranial subdural space
Mandibular canal
Myocardium of ventricle
Scapular region of back
Rhopheocytotic vesicle
Corneal corpuscle
Rotator cuff including muscles and tendons
Submucosa of anal canal
Occipital angle of parietal bone
Olivocerebellar fibers
Proximal phalanx of third toe
Ligament of diaphragm
Helper cell
Lamina propria of ethmoid sinus
Entire first left aortic arch
Abdominopelvic portion of sympathetic nervous system
Skin of glans penis
Articulations of auditory ossicles
Mucous membrane of tongue
Anterior communicating artery
Inflow tract of right ventricle
Limitans nucleus
Subcutaneous acromial bursa
Superficial flexor tendon of little finger
Membrane-coating granule, amorphous
Lateral nuclei of globus pallidus
Pancreatic veins
Entire superficial circumflex iliac vein
Stratum lemnisci of corpora quadrigemina
Radial nerve
Intervertebral disc space of twelfth thoracic vertebra
Infundibulum of Fallopian tube
Intranuclear crystal
Hindgut
Entire delphian lymph node
Supraaortic valve area
Superior anastomotic vein
Vein of head
Interlobar duct of pancreas
Superior colliculus of corpora quadrigemina
Entire lateral striate artery
Infraorbital nerve
Superior articular process of fifth thoracic vertebra
Wrist
Accessory atrioventricular bundle
Apical branch of right pulmonary artery
Osseous portion of Eustachian tube
Tunica interna of eyeball
Articular surface, metacarpal, of phalanx of hand
Small intestine serosa
Below knee region
Interlobular arteries of liver
Mastoid fontanel of skull
Lumbar lymph node
Colic lymph node
Sphincter pupillae muscle
Jugum of sphenoid bone
Lamina of eighth thoracic vertebra
Birth canal
Iliac fossa
Renal surface of adrenal gland
Joint of lumbar vertebra
Ligament of sacroiliac joint and pubic symphysis
Sinoatrial node branch of right coronary artery
Mesial surface of tooth
Obliquus capitis muscle
Inferior articular process of twelfth thoracic vertebra
Posterior intercavernous sinus
Lipid droplet
Entire juxtaintestinal lymph node
Interclavicular ligament
Both feet
Meissner's plexus
Vestibulocochlear nerve
Cricoid cartilage
Adductor hallucis muscle
Medulla oblongata fasciculus cuneatus
Right margin of heart
Zygomatic region of face
Transplanted ureter
Superior right pulmonary vein
Choroidal branches of posterior spinal artery
Glycogen vacuole
All toes
Body of right atrium
Lateral olfactory gyrus
Intervertebral foramen of second lumbar vertebra
Entire minor sublingual ducts
Periodontal tissues
Subcutaneous tissue of interdigital space of hand
Cavernous portion of internal carotid artery
Tendinous arch
Intranuclear body, granular with filamentous capsule
Corticomedullary junction of adrenal gland
Iliac tuberosity
Thenar and hypothenar spaces
Pedicle of eleventh thoracic vertebra
Peroneal artery
Shaft of phalanx of middle finger
Agranular endoplasmic reticulum, connection with other organelle
Entire subtendinous prepatellar bursa
Proper fasciculus
Crista galli
Palmar surface of middle finger
Mandibular right second premolar tooth
Brachiocephalic vein
Diaphragmatic surface of lung
Entire gastric cardiac gland (body structure)
Lateral glossoepiglottic fold
Entire left ulnar artery
Entire inferior transverse scapular ligament
Entire endocardium of right ventricle
Inguinal lymph node
Coracoid process of scapula
Cerebral meninges
Trapezoid ligament
Stratum zonale of corpora quadrigemina
Left eye
Joint of vertebral column
Marginal part of orbicularis oris muscle
Hepatic vein
Cerebellar peduncle
Left parietal lobe
Entire middle colic vein
Ascending colon
Both forearms
White matter of insula
Splenic sinusoid
Superior laryngeal vein
Arch of foot
Vein of the scala tympani
Transverse folds of palate
Embryo stage 1 structure
Capsule of metatarsophalangeal joint of fifth toe
Filaments of contractile apparatus
Intervertebral disc of eighth thoracic vertebra
Centriole
Shaft of fifth metatarsal bone
Structure of lumbar rotator muscle (body structure)
Entire external pudendal vein
Niemann-Pick cell
Posterior segment of right lobe of liver
Gravid uterus
Tendon and tendon sheath of second toe
Pelvic fascia
Corpus cavernosum of penis
Posterior intraoccipital synchondrosis
Labial veins
Merkel's tactile disc
Subtendinous iliac bursa
Tail of epididymis
Interdental papilla of gingiva
Lateral ligament of temporomandibular joint
Skin of medial surface of middle finger
All permanent teeth
Pecten ani
Lumbar vein
Lymphatics of stomach
Plantar surface of fourth toe
Structure of deep cervical lymphatics
Subclavian vein
Structure of medial cartilaginous lamina of pharyngotympanic tube (body structure)
Structure of amacrine cell of retina (body structure)
Afferent glomerular arteriole
Pulmonary ligament
Head of metacarpal bone
Coronal depression of tooth
Calcaneocuboidal ligament
Pyramid of medulla oblongata
Entire facet for fifth costal cartilage of sternum
Duodenal lumen
Subcutaneous tissue of areola
Deep branch of ulnar nerve
Posterior process of nasal septal cartilage
Lanugo hair
Left superior vena cava
Entire superior transverse scapular ligament
Gastric mucous gland
Infraclavicular lymph node
Subcutaneous tissue of lower margin of nasal septum
Ciliary muscle
Head of second metatarsal bone
Melanocyte
Posterior scrotal branches of internal pudendal artery
Iliac fascia
Right wrist
Entire tendon of index finger
Submucosa of tonsil
Genital tubercle
Left carotid sinus
Distinctive shape of mitochondrial cristae
Superficial lymphatics of thorax
Deep venous system of lower extremity
Skeletal muscle fiber, type IIb
Fascia of upper extremity
Proximal phalanx of little toe
Perforating branches of internal thoracic artery
Biparietal diameter of head
Interspinalis thoracis muscles
Right kidney
Hilum of adrenal gland
Fornix of lacrimal sac
Carunculae hymenales
Thymus
Entire appendicular vein
Thyroid tubercle
Peripheral nerve myelinated nerve fiber
Transverse arytenoid muscle
Paracentral lobule
Posterior ethmoidal nerve
Primary foot process
Ileocecal ostium
Rhomboideus cervicis muscle
Superior articular process of sixth thoracic vertebra
Duodenal ampulla
Lateral meniscus of knee joint
Base of lung
Base of phalanx of index finger
Ventral spinocerebellar tract of pons
Nucleus pulposus of intervertebral disc of eighth thoracic vertebra
Intervertebral foramen of fifth thoracic vertebra
Transplanted lung
Male
Ophthalmic nerve
Levator labii superioris muscle
Deep volar arch of radial artery
Deep dorsal sacrococcygeal ligament
Medial surface of third toe
Zygomatic nerve
Vagus nerve pharyngeal plexus
Entire costal groove of fifth rib
Merkel's cell
Middle part of cartilaginous portion of auditory canal
Entire inner surface of eleventh rib
Supraoptic region of hypothalamus
Liver
Gynecoid pelvis
Parotid lymph node
Abductor digiti minimi muscle of foot
Entire anterior commissure of aortic valve
Iliac artery
Uropod
Entire anteromedial perforating artery
Capsule of calcaneocuboidal joint
Anulus fibrosus of intervertebral disc of tenth thoracic vertebra
Sweat gland
Pigmented layer of ciliary body
Lesser sciatic notch
Pulmonic area

*/
